package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
	"time"
	"sync"
)

func parameterHandler() (string,string,int,int) {
	var ip string
	var port string
	var mask int
	var threads int
	
	flag.StringVar(&ip, "ip", "none", "Ip of network")
	flag.StringVar(&port, "p", "", "port to scan for")
	flag.IntVar(&mask, "m", 24, "network mask")
	flag.IntVar(&threads, "t", 200, "number of workers | threads")
	
	flag.Parse()

	if(mask == 0){
		ip = "0.0.0.0"
	}
	if(port == ""){
		fmt.Println("It's mandatory to select a port number")
		os.Exit(1)
	}

	
	return ip, port, mask, threads
}

func loopSelector(wg *sync.WaitGroup, ip string, port string, mask int, threads int) {
	sem := make(chan struct{}, threads)

	switch mask {
	case 24:
		ipSplited := strings.Split(ip, ".")
		for n4 := 0; n4 < 255; n4++ {
			ipSplited[3] = strconv.Itoa(n4)
			ipTarget := ipSplited[0] + "." + ipSplited[1] + "." + ipSplited[2] + "." + ipSplited[3]
			socket := ipTarget + ":" + port

			sem <- struct{}{}
			wg.Add(1)

			go func() {
				defer func() {
					<-sem
					wg.Done()
				}()

				mainDial(socket, ipTarget, port)
			}()
		}
	case 16:
		ipSplited := strings.Split(ip, ".")
		for n4:=0; n4<255; n4++ {
			for n3:=0; n3<=255; n3++ {
				ipSplited[3] = strconv.Itoa(n4)
				ipSplited[2] = strconv.Itoa(n3)
				ipTarget := ipSplited[0] + "." + ipSplited[1] + "." + ipSplited[2] + "." + ipSplited[3]
				socket := ipTarget + ":" + port

				sem <- struct{}{}
				wg.Add(1)
	
				go func() {
					defer func() {
						<-sem
						wg.Done()
					}()
	
					mainDial(socket, ipTarget, port)
				}()
			}
		}
	case 8:
		ipSplited := strings.Split(ip, ".")
		for n4:=0; n4<255; n4++ {
			for n3:=0; n3<255; n3++ {
				for n2:=0; n2<=255; n2++ {
					ipSplited[3] = strconv.Itoa(n4)
					ipSplited[2] = strconv.Itoa(n3)
					ipSplited[1] = strconv.Itoa(n2)
					ipTarget := ipSplited[0] + "." + ipSplited[1] + "." + ipSplited[2] + "." + ipSplited[3]
					socket := ipTarget + ":" + port

					sem <- struct{}{}
					wg.Add(1)
		
					go func() {
						defer func() {
							<-sem
							wg.Done()
						}()
		
						mainDial(socket, ipTarget, port)
					}()
				}
			}
		}
	case 0:
		ipSplited := strings.Split(ip, ".")
		for n4:=0; n4<=255; n4++ {
			for n3:=0; n3<=255; n3++ {
				for n2:=0; n2<=255; n2++ {
					for n1:=0; n1<255; n1++ {
						ipSplited[3] = strconv.Itoa(n4)
						ipSplited[2] = strconv.Itoa(n3)
						ipSplited[1] = strconv.Itoa(n2)
						ipSplited[0] = strconv.Itoa(n1)
						ipTarget := ipSplited[0] + "." + ipSplited[1] + "." + ipSplited[2] + "." + ipSplited[3]
						socket := ipTarget + ":" + port

						sem <- struct{}{}
						wg.Add(1)
			
						go func() {
							defer func() {
								<-sem
								wg.Done()
							}()
			
							mainDial(socket, ipTarget, port)
						}()
					}
				}
			}
		}
	}  
}

func mainDial(socket string,ipTarget string, port string) {

	dial, err := net.DialTimeout("tcp", socket, 2*time.Second )
	
	if err != nil {
		return;
	}else{
		dial.Close()
		fmt.Printf("%s has port %s open                    \n", ipTarget, port)
		return;
	}
}


func main(){
	var wg sync.WaitGroup
	ip, port, mask, threads := parameterHandler()
	loopSelector(&wg, ip, port, mask, threads)
	wg.Wait()
}