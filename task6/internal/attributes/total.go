package attributes

type COCOMOAttributes struct {
	*ComputerAttributes
	*ProgramAttributes
	*StaffAttributes
	*ProjectAttributes
}

type COCOMOAttributesConfig struct {
	ComputerAttributes *ComputerAttributesConfig
	ProgramAttributes  *ProgramAttributesConfig
	StaffAttributes    *StaffAttributesConfig
	ProjectAttributes  *ProjectAttributesConfig
}

func MapTotalAttributes(c *COCOMOAttributesConfig) *COCOMOAttributes {
	return &COCOMOAttributes{
		ComputerAttributes: MapComputerAttributes(c.ComputerAttributes),
		ProgramAttributes:  MapProgramAttributes(c.ProgramAttributes),
		StaffAttributes:    MapStaffAttributes(c.StaffAttributes),
		ProjectAttributes:  MapProjectAttributes(c.ProjectAttributes),
	}
}

func (c *COCOMOAttributes) EAF() float64 {
	return c.ProgramAttributes.CPLX * c.ProgramAttributes.Data * c.ProgramAttributes.Rely *
		c.ProjectAttributes.MODP * c.ProjectAttributes.TOOL * c.ProjectAttributes.SCED *
		c.ComputerAttributes.Time * c.ComputerAttributes.Stor * c.ComputerAttributes.Virt * c.ComputerAttributes.Turn *
		c.StaffAttributes.ACAP * c.StaffAttributes.AEXP * c.StaffAttributes.PCAP * c.StaffAttributes.VEXP * c.StaffAttributes.LEXP
}
