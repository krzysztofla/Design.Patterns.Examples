package main

// So like you see segregation of functionalities and interfaces are relevant
type Smartphone interface {
	Call(number int)
	MMS(message string)
	VideoCall(number string)
}

// For new type of phones it make sense to implement Smartphone interface
// because new type of devices has everything needed to handle
type NewFashionPhone struct {
}

func (nfp NewFashionPhone) Call(number int) {

}

func (nfp NewFashionPhone) MMS(message string) {

}

func (nfp NewFashionPhone) VideoCall(number string) {

}

// But for old fashion phones it doesn't make sense to implement everything from Smartphone interface
// because old phones doesn't have for example cameras, internet connection...
type OldFashionPhone struct {
}

// So for that kind of things the Interface Segregation Principle is ideal.
// Let's split interfaces

type VoiceCall interface {
	Call(number int)
}

type VideoCall interface {
	CallOverInternet(number int)
}

type MultimediaMessage interface {
	SentPic(message string)
}

type InterfaceForOldPhones interface {
	VoiceCall
	MultimediaMessage
}

//------
func ProducePhone(p Smartphone) *NewFashionPhone {
	return &NewFashionPhone{}
}

func mainISP() {

}
