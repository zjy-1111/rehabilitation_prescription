package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_EXIST_AUTH_FAIL          = 20005
	ERROR_EXIST_AUTH               = 20006
	ERROR_ADD_AUTH_FAIL            = 20007
	ERROR_NOT_EXIST_AUTH           = 20008
	ERROR_EDIT_AUTH_FAIL           = 20009
	ERROR_DELETE_AUTH_FAIL         = 20010

	_                        = iota
	ERROR_COUNT_PATIENT_FAIL = 10000 + iota
	ERROR_GET_PATIENTS_FAIL
	ERROR_ADD_APPOINTMENT_FAIL
	ERROR_CHECK_EXIST_APPOINTMENT_FAIL
	ERROR_NOT_EXIST_APPOINTMENT
	ERROR_EDIT_APPOINTMENT_FAIL
	ERROR_DELETE_APPOINTMENT_FAIL

	ERROR_COUNT_PRESCRIPTION_FAIL
	ERROR_GET_PRESCRIPTIONS_FAIL
	ERROR_ADD_PRESCRIPTION_FAIL
	ERROR_CHECK_EXIST_PRESCRIPTION_FAIL
	ERROR_NOT_EXIST_PRESCRIPTION
	ERROR_EDIT_PRESCRIPTION_FAIL
	ERROR_DELETE_PRESCRIPTION_FAIL

	ERROR_COUNT_PATIENTREPORT_FAIL
	ERROR_GET_PATIENTREPORTS_FAIL
	ERROR_ADD_PATIENTREPORT_FAIL
	ERROR_CHECK_EXIST_PATIENTREPORT_FAIL
	ERROR_NOT_EXIST_PATIENTREPORT
	ERROR_EDIT_PATIENTREPORT_FAIL
	ERROR_DELETE_PATIENTREPORT_FAIL

	ERROR_COUNT_REPORTPRESCRIPTION_FAIL
	ERROR_GET_REPORTPRESCRIPTIONS_FAIL
	ERROR_ADD_REPORTPRESCRIPTION_FAIL
	ERROR_CHECK_EXIST_REPORTPRESCRIPTION_FAIL
	ERROR_NOT_EXIST_REPORTPRESCRIPTION
	ERROR_EDIT_REPORTPRESCRIPTION_FAIL
	ERROR_DELETE_REPORTPRESCRIPTION_FAIL

	ERROR_COUNT_TRAININGVIDEO_FAIL
	ERROR_GET_TRAININGVIDEOS_FAIL
	ERROR_ADD_TRAININGVIDEO_FAIL
	ERROR_CHECK_EXIST_TRAININGVIDEO_FAIL
	ERROR_NOT_EXIST_TRAININGVIDEO
	ERROR_EDIT_TRAININGVIDEO_FAIL
	ERROR_DELETE_TRAININGVIDEO_FAIL

	ERROR_COUNT_PRESCRIPTIONVIDEO_FAIL
	ERROR_GET_PRESCRIPTIONVIDEOS_FAIL
	ERROR_ADD_PRESCRIPTIONVIDEO_FAIL
	ERROR_CHECK_EXIST_PRESCRIPTIONVIDEO_FAIL
	ERROR_NOT_EXIST_PRESCRIPTIONVIDEO
	ERROR_EDIT_PRESCRIPTIONVIDEO_FAIL
	ERROR_DELETE_PRESCRIPTIONVIDEO_FAIL
)
