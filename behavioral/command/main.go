package main

func main() {

	patient1 := &Patient{
		name:              "Billy H.",
		registrationDone:  false,
		doctorCheckUpDone: false,
		medicineDone:      false,
		paymentDone:       false,
	}

	var reception = Reception{&Doctor{next: &Medical{next: &Cashier{nil}}}}

	reception.execute(patient1)

	//Second patient that already registered

	patient2 := &Patient{
		name:              "Ricardo M.",
		registrationDone:  true,
		doctorCheckUpDone: false,
		medicineDone:      false,
		paymentDone:       false,
	}

	reception.execute(patient2)
}
