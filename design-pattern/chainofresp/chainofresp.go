package chainofresp

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
}

type department interface {
	execute(*patient)
	setNext(department)
}

// Reception
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {

	if p.registrationDone {
		r.next.execute(p)
		return
	}

	p.registrationDone = true
	r.next.execute(p)
}

// Doctor
type doctor struct {
	next department
}

func (r *doctor) execute(p *patient) {

	if p.doctorCheckUpDone {
		r.next.execute(p)
		return
	}

	p.doctorCheckUpDone = true
	r.next.execute(p)
}

// Medicine
type medicine struct {
	next department
}

func (r *medicine) execute(p *patient) {

	if p.medicineDone {
		r.next.execute(p)
		return
	}

	p.medicineDone = true
	r.next.execute(p)
}
