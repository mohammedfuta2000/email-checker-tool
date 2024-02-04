package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main()  {
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
	if err:=scanner.Err(); err!=nil{
		log.Fatal("Error: Could Not read from input: %v\n",err)
	}
}

func checkDomain(domain string)  {
	var hasMX,hasDMARC,hasSPF bool
	var sprRecord,dmarcRecord string

	mxRecords,err:= net.LookupMX(domain)
	if err!=nil{
		log.Printf(err.Error())
	}
	if len(mxRecords)>0 {
		hasMX=true
	}

	txtRecords,err:=net.LookupTXT(domain)
	if err!=nil{
		log.Printf(err.Error())
	}

	for _, record:= range txtRecords{
		if strings.HasPrefix(record,"v=spf1"){
			hasSPF=true
			sprRecord=record
			break
		}

	}

	dmarcRecords,err:= net.LookupTXT("_dmarc."+domain)
	if err!=nil{
		fmt.Printf(err.Error())
	}
	for _, record:= range dmarcRecords{
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC=true
			dmarcRecord=record
			break
		}

	}
	fmt.Printf("Domain: %v,\nhasMX: %v,\nhasSPF: %v,\nsprRecord: %v,\nhasDMARC: %v,\nhasDMARC: %v,\n",domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord)

}
