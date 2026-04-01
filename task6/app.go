package main

import (
	"context"
	"task6/internal/attributes"
	"task6/internal/chars"
	"task6/internal/cocomo"
	"task6/internal/distribution"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type CalculateInput struct {
	Mode          string  `json:"mode"`
	Kloc          float64 `json:"kloc"`
	BudgetPerWork float64 `json:"budgetPerWork"`

	ComputerAttributes *ComputerAttrInput `json:"computerAttributes"`
	ProgramAttributes  *ProgramAttrInput  `json:"programAttributes"`
	StaffAttributes    *StaffAttrInput    `json:"staffAttributes"`
	ProjectAttributes  *ProjectAttrInput  `json:"projectAttributes"`
}

type ComputerAttrInput struct {
	Time attributes.AttributeLevel `json:"time"`
	Stor attributes.AttributeLevel `json:"stor"`
	Virt attributes.AttributeLevel `json:"virt"`
	Turn attributes.AttributeLevel `json:"turn"`
}

type ProgramAttrInput struct {
	Rely attributes.AttributeLevel `json:"rely"`
	Data attributes.AttributeLevel `json:"data"`
	CPLX attributes.AttributeLevel `json:"cplx"`
}

type StaffAttrInput struct {
	ACAP attributes.AttributeLevel `json:"acap"`
	AEXP attributes.AttributeLevel `json:"aexp"`
	PCAP attributes.AttributeLevel `json:"pcap"`
	VEXP attributes.AttributeLevel `json:"vexp"`
	LEXP attributes.AttributeLevel `json:"lexp"`
}

type ProjectAttrInput struct {
	MODP attributes.AttributeLevel `json:"modp"`
	TOOL attributes.AttributeLevel `json:"tool"`
	SCED attributes.AttributeLevel `json:"sced"`
}

type CalculateOutput struct {
	Work         float64           `json:"work"`
	Time         float64           `json:"time"`
	WorkChars    WorkCharsOutput   `json:"workChars"`
	TimeChars    TimeCharsOutput   `json:"timeChars"`
	BudgetChars  BudgetCharsOutput `json:"budgetChars"`
	WorkCoeffs   WorkCharsOutput   `json:"workCoeffs"`
	TimeCoeffs   TimeCharsOutput   `json:"timeCoeffs"`
	BudgetCoeffs BudgetCharsOutput `json:"budgetCoeffs"`
	BudgetWork   BudgetCharsOutput `json:"budgetWork"`
}

type WorkCharsOutput struct {
	Planning           float64 `json:"planning"`
	Design             float64 `json:"design"`
	Construction       float64 `json:"construction"`
	CodingTesting      float64 `json:"codingTesting"`
	IntegrationTesting float64 `json:"integrationTesting"`
}

type WorkCharsInput struct {
	Planning           float64 `json:"planning"`
	Design             float64 `json:"design"`
	Construction       float64 `json:"construction"`
	CodingTesting      float64 `json:"codingTesting"`
	IntegrationTesting float64 `json:"integrationTesting"`
}

type TimeCharsOutput struct {
	Planning           float64 `json:"planning"`
	Design             float64 `json:"design"`
	Construction       float64 `json:"construction"`
	CodingTesting      float64 `json:"codingTesting"`
	IntegrationTesting float64 `json:"integrationTesting"`
}

type TimeCharsInput struct {
	Planning           float64 `json:"planning"`
	Design             float64 `json:"design"`
	Construction       float64 `json:"construction"`
	CodingTesting      float64 `json:"codingTesting"`
	IntegrationTesting float64 `json:"integrationTesting"`
}

type BudgetCharsOutput struct {
	Analysis                float64 `json:"analysis"`
	Design                  float64 `json:"design"`
	Programming             float64 `json:"programming"`
	TestPlanning            float64 `json:"testPlanning"`
	Verification            float64 `json:"verification"`
	Office                  float64 `json:"office"`
	ConfigurationManagement float64 `json:"configManagement"`
	Documentation           float64 `json:"documentation"`
}

type DistributionInput struct {
	Work      float64        `json:"work"`
	Time      float64        `json:"time"`
	WorkChars WorkCharsInput `json:"workChars"`
	TimeChars TimeCharsInput `json:"timeChars"`
}

func (a *App) CalculateDistribution(input DistributionInput) []int {
	workChars := &chars.WorkCharacteristics{
		Planning:           input.WorkChars.Planning,
		Design:             input.WorkChars.Design,
		Construction:       input.WorkChars.Construction,
		CodingTesting:      input.WorkChars.CodingTesting,
		IntegrationTesting: input.WorkChars.IntegrationTesting,
	}
	timeChars := &chars.TimeCharacteristics{
		Planning:           input.TimeChars.Planning,
		Design:             input.TimeChars.Design,
		Construction:       input.TimeChars.Construction,
		CodingTesting:      input.TimeChars.CodingTesting,
		IntegrationTesting: input.TimeChars.IntegrationTesting,
	}
	calc := distribution.NewHumanDistributionCalculator(workChars, timeChars)
	return calc.Calculate()
}

func (a *App) Calculate(input CalculateInput) CalculateOutput {
	cocomoMode := cocomo.COCOMOMode(input.Mode)

	attrConfig := &attributes.COCOMOAttributesConfig{
		ComputerAttributes: &attributes.ComputerAttributesConfig{
			Time: input.ComputerAttributes.Time,
			Stor: input.ComputerAttributes.Stor,
			Virt: input.ComputerAttributes.Virt,
			Turn: input.ComputerAttributes.Turn,
		},
		ProgramAttributes: &attributes.ProgramAttributesConfig{
			Rely: input.ProgramAttributes.Rely,
			Data: input.ProgramAttributes.Data,
			CPLX: input.ProgramAttributes.CPLX,
		},
		StaffAttributes: &attributes.StaffAttributesConfig{
			ACAP: input.StaffAttributes.ACAP,
			AEXP: input.StaffAttributes.AEXP,
			PCAP: input.StaffAttributes.PCAP,
			VEXP: input.StaffAttributes.VEXP,
			LEXP: input.StaffAttributes.LEXP,
		},
		ProjectAttributes: &attributes.ProjectAttributesConfig{
			MODP: input.ProjectAttributes.MODP,
			TOOL: input.ProjectAttributes.TOOL,
			SCED: input.ProjectAttributes.SCED,
		},
	}

	attrs := attributes.MapTotalAttributes(attrConfig)
	calc := cocomo.NewCOCOMOCalculator(input.Kloc, attrs)
	calc.SetMode(cocomoMode)

	workResult, timeResult := calc.Calculate()

	workCharsResult := chars.CalculateWorkCharacteristics(workResult)
	timeCharsResult := chars.CalculateTimeCharacteristics(timeResult)
	budgetChars := chars.CalculateBudgetCharacteristics(workResult, input.BudgetPerWork)
	budgetWorkChars := chars.CalculateBudgetWorkCharacteristics(workResult)

	return CalculateOutput{
		Work: workResult,
		Time: timeResult,
		WorkChars: WorkCharsOutput{
			Planning:           workCharsResult.Planning,
			Design:             workCharsResult.Design,
			Construction:       workCharsResult.Construction,
			CodingTesting:      workCharsResult.CodingTesting,
			IntegrationTesting: workCharsResult.IntegrationTesting,
		},
		TimeChars: TimeCharsOutput{
			Planning:           timeCharsResult.Planning,
			Design:             timeCharsResult.Design,
			Construction:       timeCharsResult.Construction,
			CodingTesting:      timeCharsResult.CodingTesting,
			IntegrationTesting: timeCharsResult.IntegrationTesting,
		},
		BudgetChars: BudgetCharsOutput{
			Analysis:                budgetChars.Analysis,
			Design:                  budgetChars.Design,
			Programming:             budgetChars.Programming,
			TestPlanning:            budgetChars.TestPlanning,
			Verification:            budgetChars.Verification,
			Office:                  budgetChars.Office,
			ConfigurationManagement: budgetChars.ConfigurationManagement,
			Documentation:           budgetChars.Documentation,
		},
		WorkCoeffs: WorkCharsOutput{
			Planning:           chars.WorkCharacteristicsCoefficients.Planning,
			Design:             chars.WorkCharacteristicsCoefficients.Design,
			Construction:       chars.WorkCharacteristicsCoefficients.Construction,
			CodingTesting:      chars.WorkCharacteristicsCoefficients.CodingTesting,
			IntegrationTesting: chars.WorkCharacteristicsCoefficients.IntegrationTesting,
		},
		TimeCoeffs: TimeCharsOutput{
			Planning:           chars.TimeCharacteristicsCoefficients.Planning,
			Design:             chars.TimeCharacteristicsCoefficients.Design,
			Construction:       chars.TimeCharacteristicsCoefficients.Construction,
			CodingTesting:      chars.TimeCharacteristicsCoefficients.CodingTesting,
			IntegrationTesting: chars.TimeCharacteristicsCoefficients.IntegrationTesting,
		},
		BudgetCoeffs: BudgetCharsOutput{
			Analysis:                chars.BudgetCharacteristicsCoefficients.Analysis,
			Design:                  chars.BudgetCharacteristicsCoefficients.Design,
			Programming:             chars.BudgetCharacteristicsCoefficients.Programming,
			TestPlanning:            chars.BudgetCharacteristicsCoefficients.TestPlanning,
			Verification:            chars.BudgetCharacteristicsCoefficients.Verification,
			Office:                  chars.BudgetCharacteristicsCoefficients.Office,
			ConfigurationManagement: chars.BudgetCharacteristicsCoefficients.ConfigurationManagement,
			Documentation:           chars.BudgetCharacteristicsCoefficients.Documentation,
		},
		BudgetWork: BudgetCharsOutput{
			Analysis:                budgetWorkChars.Analysis,
			Design:                  budgetWorkChars.Design,
			Programming:             budgetWorkChars.Programming,
			TestPlanning:            budgetWorkChars.TestPlanning,
			Verification:            budgetWorkChars.Verification,
			Office:                  budgetWorkChars.Office,
			ConfigurationManagement: budgetWorkChars.ConfigurationManagement,
			Documentation:           budgetWorkChars.Documentation,
		},
	}
}
