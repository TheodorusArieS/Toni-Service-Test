package query

const (
	InsertStatusDefault = `
	INSERT INTO statuses(name,description)
	VALUES	('Register','User Register'),
			('Completed','User Register Complete'),
			('Pending','Transaction Status Pending'),
			('Success','Transaction Status Success')
	`

	QueryCreateOTP = `
	INSERT INTO otp_verification(phone,otp,type_id)
	VALUES (?,?,?)
	`

	QueryValidateOTP = `
	SELECT id
	FROM otp_verification
	WHERE phone = ? AND otp = ? AND type_id = ? 
	`

	QueryCreateNewUser = `
	INSERT INTO users(phone,pin_number,status_id)
	VALUES (?,?,?)
	`

	QueryDeleteValidateOTP =`
	DELETE FROM otp_verification
	WHERE id = ?
	`	

	QueryLogin =`
	SELECT id,phone,pin_number
	FROM users
	WHERE phone = ?
	
	`


)
