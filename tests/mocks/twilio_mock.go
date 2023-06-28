package mocks


type TwilioMock struct {
	TimesInvoked int
}

func (t *TwilioMock) SendSMS(toPhoneNumber string, fromPhoneNumber string, message string) error {
	t.registerInvokation()

	return nil
}

func (t *TwilioMock) registerInvokation() {
    t.TimesInvoked++
}
