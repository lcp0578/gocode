package main

import (
	"fmt"
	"gitee.com/lcp0578/go_kit/crc"

)

func main(){
	data := []byte{
		0x01,
		0x30,
		0x31,
		0x32,
		0x33,
		0x34,
		0x35,
		0x36,
		0x37,
		0x32,
		0x38,
		0x30,
		0x31,
		0x38,
		0x38,
		0x38,
		0x38,
		0x38,
		0x31,
		0x30,
		0x02,
		0x30,
		0x30,
		0x30,
		0x30,
		0x31,
		0x38,
		0x30,
		0x38,
		0x31,
		0x31,
		0x31,
		0x30,
		0x32,
		0x36,
		0x33,
		0x39,
		//0x1b,
		//0x65,
		//0x33,
		//0x62,
	}

	crcData := crc.CalculateCRC(crc.CRC16, data)
	//checksum := crc16.ChecksumIBM(data)
	crcHex := fmt.Sprintf("%X", crcData)
	//buff := new(bytes.Buffer)
	//err := binary.Write(buff, binary.LittleEndian, uint16(crcData))
	//if err != nil  {
	//	fmt.Println(err)
	//}
	//fmt.Println(buff.Bytes())
	//crcHex := strconv.FormatUint(crcData, 16)
	//fmt.Print(crcHex)
	crcVal := []byte(crcHex)
	fmt.Println(crcVal)
	//
	//crcHex = strconv.FormatUint(uint64(crc16.ChecksumCCITT(data)), 16)
	//fmt.Print(crcHex)
	//crcVal = []byte(crcHex)
	//fmt.Println(crcVal)
	//
	//crcHex = strconv.FormatUint(uint64(crc16.ChecksumCCITTFalse(data)), 16)
	//fmt.Print(crcHex)
	//crcVal = []byte(crcHex)
	//fmt.Println(crcVal)
	//
	//crcHex = strconv.FormatUint(uint64(crc16.ChecksumSCSI(data)), 16)
	//fmt.Print(crcHex)
	//crcVal = []byte(crcHex)
	//fmt.Println(crcVal)
	//
	//crcHex = strconv.FormatUint(uint64(crc16.ChecksumMBus(data)), 16)
	//fmt.Print(crcHex)
	//crcVal = []byte(crcHex)
	//fmt.Println(crcVal)
}
