// +build httpcontent all

package httpcontent

import (
	"encoding/xml"
//	"fmt"
)


type XMLRequest struct {
	XMLName xml.Name `xml:""`
	Target string `xml:",attr"`

	POS struct {
		Source struct {
			TerminalID string `xml:",attr"`
		} `xml:"Source"`
	} `xml:"POS"`

}



func (t *HttpContentTunnel) xmlParse(xmldoc string) (string,string,string, error){

	var xr XMLRequest

	err := xml.Unmarshal([]byte(xmldoc), &xr)
	if err != nil {
		//fmt.Println(err)
		return "","","", err
	}

/*	fmt.Println(xr.XMLName.Local)
	fmt.Println(xr.Target)
	fmt.Println(xr.POS.Source.TerminalID)*/

	return xr.XMLName.Local, xr.Target, xr.POS.Source.TerminalID, nil

}

