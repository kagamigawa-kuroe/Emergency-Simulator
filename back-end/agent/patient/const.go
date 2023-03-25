package patient

type patient_status int8

const (
	Waiting_for_nurse      patient_status = 1
	Being_judged_by_nurse  patient_status = 2
	Waiting_for_register   patient_status = 3
	Waiting_for_treat      patient_status = 4
	Being_treated_now      patient_status = 5
	Being_treated_urgently patient_status = 6
	Finish                 patient_status = 7
)
