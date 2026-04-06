package functional

type Language string

var (
	AssemblyLanguage Language = "Assembly"
	C                Language = "C"
	Cobol            Language = "Cobol"
	Fortran          Language = "Fortran"
	Pascal           Language = "Pascal"
	CPlusPlus        Language = "C++"
	CSharp           Language = "C#"
	Java             Language = "Java"
	Ada95            Language = "Ada95"
	VisualBasic      Language = "VisualBasic"
	VisualCpp        Language = "VisualCpp"
	Delphi           Language = "Delphi"
	Perl             Language = "Perl"
	Prolog           Language = "Prolog"
	SQL              Language = "SQL"
	Lisp             Language = "Lisp"
)

var (
	// lines of code per functional point
	locPerFP = map[Language]uint{
		AssemblyLanguage: 320,
		C:                128,
		Cobol:            106,
		Fortran:          105,
		Pascal:           91,
		CPlusPlus:        53,
		CSharp:           53,
		Java:             53,
		Ada95:            49,
		VisualBasic:      32,
		VisualCpp:        34,
		Delphi:           18,
		Perl:             21,
		Prolog:           64,
		SQL:              13,
		Lisp:             64,
	}
)
