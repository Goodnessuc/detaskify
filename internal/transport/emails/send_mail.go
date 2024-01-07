package email

// SendVerifyAccount sends a welcome message to a new user
func (e *EmailService) SendVerifyAccount(username, url, recipient string) error {
	message := Message{
		Header:       "Tiventer",
		Username:     username,
		Introduction: "Hey " + username + ", Welcome to Tiventer",
		Content:      "Thank you for signing up on Tiventer.\n \n Please click the link below to activate your account.",
		URL:          url,
		ActionTitle:  "Verify Account",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Welcome to Tiventer", body)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmailService) SendCreateAccountEmail(username, url, recipient string) error {
	message := Message{
		Header:       "Welcome to Tiventer",
		Username:     username,
		Introduction: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Content:      "We are excited to have you on board. Start exploring our features and services.",
		URL:          url,
		ActionTitle:  "Get Started",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Your New Tiventer Account", body)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendResetAccountEmail(username, url, recipient string) error {
	message := Message{
		Header:       "Account Reset",
		Username:     username,
		Introduction: "You have requested to reset your account.",
		Content:      "Please click the link below to proceed with resetting your account.",
		URL:          url,
		ActionTitle:  "Reset Account",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Account Reset Request", body)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendChangePasswordEmail(username, recipient, url string) error {
	message := Message{
		Header:       "Password Change Confirmation",
		Username:     username,
		Introduction: "Your password has been successfully changed.",
		Content:      "If you did not initiate this change, please contact our support team immediately.",
		URL:          url,
		ActionTitle:  "Contact Support",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Password Change Notification", body)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendVerificationSuccessEmail(username, recipient, url string) error {
	message := Message{
		Header:       "Account Verified",
		Username:     username,
		Introduction: "Congratulations! Your account has been successfully verified.",
		Content:      "You can now enjoy all the features of our service.",
		URL:          url,
		ActionTitle:  "Visit the Homepage",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Account Verification Successful", body)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendResetPasswordConfirmation(username, recipient, url string) error {
	message := Message{
		Header:       "Password Reset Confirmation",
		Username:     username,
		Introduction: "Your password has been successfully reset.",
		Content:      "You can now log in with your new password.\n\n If you did not initiate this change, please contact our support team immediately.",
		URL:          url,
		ActionTitle:  "Contact Support",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Password Reset Successful", body)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendOTPEmail(username, otp, recipient, url string) error {
	message := Message{
		Header:       "One-Time Password (OTP) for Login",
		Username:     username,
		Introduction: "Your OTP for login is below.",
		Content:      "Use this code to complete your login",
		URL:          "",
		ActionTitle:  otp,
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Your Login OTP", body)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendDeleteAccountConfirmation(username, url, recipient string) error {
	message := Message{
		Header:       "Confirm Account Deletion",
		Username:     username,
		Introduction: "You have requested to delete your account.",
		Content:      "Please click the link below to confirm your account deletion.",
		URL:          url,
		ActionTitle:  "Confirm Deletion",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Account Deletion Request", body)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendSecurityAlertEmail(username, details, recipient, url string) error {
	message := Message{
		Header:       "Security Alert: New Login Attempt",
		Username:     username,
		Introduction: "A new login to your account was detected.",
		Content:      "Details: " + details + ". If this was not you, please contact our support team immediately.",
		URL:          url,
		ActionTitle:  "Contact Support",
	}
	body, err := e.LoadEmail(message)
	if err != nil {
		return err
	}
	err = e.SendEmail(recipient, "Security Alert", body)
	if err != nil {
		return err
	}
	return nil
}
