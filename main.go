package main

import (
	"github.com/miekg/dns"
	"log"
	"strings"
	"net"
)

type MyHandler struct {

}

func (handler *MyHandler) ServeDNS(w dns.ResponseWriter, msg *dns.Msg)  {
	response := new(dns.Msg)
	response.SetReply(msg)

	for _, question := range msg.Question {
		if strings.HasSuffix(question.Name, ".test.") {
			rr := &dns.A{
				Hdr: dns.RR_Header{
					Name: question.Name,
					Rrtype: dns.TypeA,
					Class: dns.ClassINET,
					Ttl: 60,
				},
				A: net.ParseIP("127.0.0.1"),
			}
			response.Answer = append(response.Answer, rr)
			continue
		}
		if strings.HasSuffix(question.Name, ".xip.io.")  {
			handler.handleDynamic(response, question.Name, "xip.io.")
			continue
		}
		if strings.HasSuffix(question.Name, ".nip.io.")  {
			handler.handleDynamic(response, question.Name, "nip.io.")
			continue
		}
	}

	w.WriteMsg(response)
}
func (handler *MyHandler) handleDynamic(response *dns.Msg, name string, suffix string) {
	sa := strings.Split(suffix, ".")
	na := strings.Split(name, ".")

	address := strings.Join(na[len(na)-len(sa)-4:len(na)-len(sa)], ".")

	rr := &dns.A{
		Hdr: dns.RR_Header{
			Name: name,
			Rrtype: dns.TypeA,
			Class: dns.ClassINET,
			Ttl: 60,
		},
		A: net.ParseIP(address),
	}

	response.Answer = append(response.Answer, rr)
}

func main() {
	handler := MyHandler{}

	server := dns.Server{
		Addr: "127.0.0.1:5300",
		Net: "udp",
		Handler: &handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}