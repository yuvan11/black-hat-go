package rpc

import (
	"bytes"
	"fmt"
	"net/http"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

func New(host, user, pass string) (*Metasploit, error) {

	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}
	fmt.Println("MSF", msf)
	if err := msf.Login(); err != nil {
		return nil, err
	}

	return msf, nil

}

type loginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logoutReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logoutRes struct {
	Result string `msgpack:"result"`
}
type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListRes struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

func (msf *Metasploit) send(req interface{}, res interface{}) error {

	fmt.Println("Reg", req)

	buff := new(bytes.Buffer)
	msgpack.NewEncoder(buff).Encode(req)

	dest := fmt.Sprintf("http://%s/api", msf.host)

	resp, err := http.Post(dest, "binary/message-pack", buff)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := msgpack.NewDecoder(resp.Body).Decode(&res); err != nil {

		return err
	}
	return nil

}

func (msf *Metasploit) Login() error {

	loginreq := &loginReq{

		Method:   "auth.login",
		Username: msf.user,
		Password: msf.pass,
	}

	var res loginRes

	if err := msf.send(loginreq, &res); err != nil {

		return err
	}
	fmt.Println("resp", res)
	msf.token = res.Token
	fmt.Println("Login token", msf.token)
	return nil

}

func (msf *Metasploit) Logout() error {

	logoutreq := logoutReq{
		Method:      "auth.logout",
		Token:       msf.token,
		LogoutToken: msf.token,
	}

	var res logoutRes

	if err := msf.send(logoutreq, &res); err != nil {
		return nil
	}
	msf.token = ""
	return nil

}

func (msf *Metasploit) SessionList() (map[uint32]SessionListRes, error) {

	req := &sessionListReq{

		Method: "session.list",
		Token:  msf.token,
	}
	fmt.Println("TOken", msf.token)

	res := make(map[uint32]SessionListRes)

	if err := msf.send(req, &res); err != nil {
		fmt.Println("ses err", err)
		return nil, err
	}

	// setting up session id and session value
	for id, session := range res {
		session.ID = id
		res[id] = session
	}

	return res, nil

}
