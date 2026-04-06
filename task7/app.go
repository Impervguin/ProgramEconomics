package main

import (
	"context"
	"fmt"
	"task7/internal/architecture"
	"task7/internal/common/team"
	"task7/internal/composition"
	"task7/internal/functional"
)

// App struct
type App struct {
	ctx                 context.Context
	functionalService   functional.FunctionalPointsService
	compositionService  composition.CompositionService
	architectureService architecture.ArchitectureService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		functionalService:   functional.FunctionalPointsService{},
		compositionService:  composition.CompositionService{},
		architectureService: architecture.ArchitectureService{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// FunctionalPointsDto for functional points calculation
type FunctionalPointsDto struct {
	Params    functional.ParamConfig `json:"params"`
	Languages map[string]float64     `json:"languages"`
	Elements  []ElementOptionDto     `json:"elements"`
}

// FunctionalPointsResult for result
type FunctionalPointsResult struct {
	Points          float64 `json:"points"`
	CorrectedPoints float64 `json:"correctedPoints"`
	LoC             float64 `json:"loc"`
}

// getElementByName returns the FunctionalElement by name
func getElementByName(name string) functional.FunctionalElement {
	switch name {
	case "externalInput":
		return &functional.ExternalInputElement
	case "externalOutput":
		return &functional.ExternalOutputElement
	case "externalQuery":
		return &functional.ExternalQueryElement
	case "internalLogicalFile":
		return &functional.InternalLogicalFileElement
	case "externalInterfaceFile":
		return &functional.ExternalInterfaceFileElement
	default:
		return nil
	}
}

// getLevelByName returns the ElementLevel by name
func getLevelByName(name string) functional.ElementLevel {
	switch name {
	case "Low":
		return functional.Low
	case "Normal":
		return functional.Normal
	case "High":
		return functional.High
	default:
		return functional.Low
	}
}

// CalculateFunctionalPoints calculates functional points
func (a *App) CalculateFunctionalPoints(dto FunctionalPointsDto) FunctionalPointsResult {
	fmt.Println("Calculating functional points...")
	elements := make([]functional.ElementOption, len(dto.Elements))
	fmt.Println(len(dto.Elements))
	for i, e := range dto.Elements {
		elements[i] = functional.ElementOption{
			Element: getElementByName(e.Element),
			Count:   e.Count,
			Level:   getLevelByName(e.Level),
		}
	}

	fp := a.functionalService.CalculatePoints(elements)

	cfp := a.functionalService.CorrectedPoints(elements,
		functional.WithDataPassing(dto.Params.DataPassing),
		functional.WithDistributedProcessing(dto.Params.DistributedProcessing),
		functional.WithEfficiency(dto.Params.Efficiency),
		functional.WithExpluatationRequirements(dto.Params.ExpluatationRequirements),
		functional.WithTransaccrionFrequency(dto.Params.TransacrionFrequency),
		functional.WithFastIO(dto.Params.FastIO),
		functional.WithUserEffectiveness(dto.Params.UserEffectiveness),
		functional.WithUpdateEfficiency(dto.Params.UpdateEfficiency),
		functional.WithProcessingComplexity(dto.Params.ProcessingComplexity),
		functional.WithReusability(dto.Params.Reusability),
		functional.WithInstallationLightness(dto.Params.InstallationLightness),
		functional.WithExpluatationLightness(dto.Params.ExpluatationLightness),
		functional.WithPlatformDiversity(dto.Params.PlatformDiversity),
		functional.WithMaintainability(dto.Params.Maintainability),
	)

	var totalLoc float64
	var totalPercent float64

	for lang, percent := range dto.Languages {
		fmt.Println(lang, percent)
		totalPercent += percent
		totalLoc += a.functionalService.FPtoLoC(fp*percent/100, functional.Language(lang))
	}
	if totalPercent != 100 {
		// error, but for now return as is
	}

	return FunctionalPointsResult{
		Points:          fp,
		CorrectedPoints: cfp,
		LoC:             totalLoc,
	}
}

// Conversion functions

func getObjectiveElement(name string) composition.ObjectiveElement {
	switch name {
	case "ViewElement":
		return &composition.ViewElement
	case "ReportElement":
		return &composition.ReportElement
	default:
		return nil
	}
}

func getProductivityLevel(name string) composition.TeamProductivityLevel {
	switch name {
	case "VeryLowProductivity":
		return composition.VeryLowProductivity
	case "LowProductivity":
		return composition.LowProductivity
	case "MediumProductivity":
		return composition.MediumProductivity
	case "HighProductivity":
		return composition.HighProductivity
	case "VeryHighProductivity":
		return composition.VeryHighProductivity
	default:
		return composition.MediumProductivity
	}
}

func getAttributeLevel(name string) architecture.AttributeLevel {
	switch name {
	case "ExtraLow":
		return architecture.ExtraLow
	case "VeryLow":
		return architecture.VeryLow
	case "Low":
		return architecture.Low
	case "Nominal":
		return architecture.Nominal
	case "High":
		return architecture.High
	case "VeryHigh":
		return architecture.VeryHigh
	case "ExtraHigh":
		return architecture.ExtraHigh
	default:
		return architecture.Nominal
	}
}

func getTeamLevel(name string) team.TeamTeamParameterLevel {
	switch name {
	case "SeverelyHindered":
		return team.SeverelyHindered
	case "SomewhatHindered":
		return team.SomewhatHindered
	case "SomeCohesion":
		return team.SomeCohesion
	case "IncreasedCohesion":
		return team.IncreasedCohesion
	case "HighCohesion":
		return team.HighCohesion
	case "UnifiedWhole":
		return team.UnifiedWhole
	default:
		return team.UnifiedWhole
	}
}

func getPmatLevel(name string) team.PmatTeamParameterLevel {
	switch name {
	case "CMMLevel1":
		return team.CMMLevel1
	case "CMMLevel1Plus":
		return team.CMMLevel1Plus
	case "CMMLevel2":
		return team.CMMLevel2
	case "CMMLevel3":
		return team.CMMLevel3
	case "CMMLevel4":
		return team.CMMLevel4
	case "CMMLevel5":
		return team.CMMLevel5
	default:
		return team.CMMLevel5
	}
}

func getPrecLevel(name string) team.PrecTeamParameterLevel {
	switch name {
	case "NewChaoticProject":
		return team.NewChaoticProject
	case "NewProject":
		return team.NewProject
	case "SomeExperienceProject":
		return team.SomeExperienceProject
	case "CommonExperienceProject":
		return team.CommonExperienceProject
	case "ExperiencedProject":
		return team.ExperiencedProject
	case "FullyKnownProject":
		return team.FullyKnownProject
	default:
		return team.FullyKnownProject
	}
}

func getReslLevel(name string) team.ReslTeamParameterLevel {
	switch name {
	case "Little20":
		return team.Little20
	case "Some40":
		return team.Some40
	case "Frequent60":
		return team.Frequent60
	case "Overall75":
		return team.Overall75
	case "AlmostFull90":
		return team.AlmostFull90
	case "Full100":
		return team.Full100
	default:
		return team.Full100
	}
}

func getFlexLevel(name string) team.FlexTeamParameterLevel {
	switch name {
	case "RigidProcess":
		return team.RigidProcess
	case "RandomRelaxations":
		return team.RandomRelaxations
	case "SomeRelaxations":
		return team.SomeRelaxations
	case "MostlyAgreedProcess":
		return team.MostlyAgreedProcess
	case "SomeAgreement":
		return team.SomeAgreement
	case "GeneralGoalsOnly":
		return team.GeneralGoalsOnly
	default:
		return team.GeneralGoalsOnly
	}
}

func convertTeamConfig(tc map[string]string) team.TeamConfig {
	return team.TeamConfig{
		TeamParameterLevel: getTeamLevel(tc["TeamParameterLevel"]),
		PmatParameterLevel: getPmatLevel(tc["PmatParameterLevel"]),
		PrecParameterLevel: getPrecLevel(tc["PrecParameterLevel"]),
		ReslParameterLevel: getReslLevel(tc["ReslParameterLevel"]),
		FlexParameterLevel: getFlexLevel(tc["FlexParameterLevel"]),
	}
}

// Frontend DTOs

type ElementOptionDto struct {
	Element string `json:"element"`
	Count   uint   `json:"count"`
	Level   string `json:"level"`
}

type CompositionModelDtoFrontend struct {
	Elements          []ElementOptionDto `json:"Elements"`
	ModuleCount       uint               `json:"ModuleCount"`
	ReuseParam        float64            `json:"ReuseParam"`
	ProductivityLevel string             `json:"ProductivityLevel"`
	AveragePay        float64            `json:"AveragePay"`
	TeamConfig        map[string]string  `json:"TeamConfig"`
}

type ArchitectureModelDtoFrontend struct {
	ArchitectureAttributes map[string]string `json:"ArchitectureAttributes"`
	TeamConfig             map[string]string `json:"TeamConfig"`
	Kloc                   float64           `json:"Kloc"`
	AveragePay             float64           `json:"AveragePay"`
}

// CalculateCompositionModel calculates composition model
func (a *App) CalculateCompositionModel(dto CompositionModelDtoFrontend) composition.CompositionModelResult {
	// Convert
	elements := make([]composition.ElementOption, len(dto.Elements))
	for i, e := range dto.Elements {
		elements[i] = composition.ElementOption{
			Element: getObjectiveElement(e.Element),
			Count:   e.Count,
			Level:   composition.ElementLevel(e.Level),
		}
	}
	serviceDto := composition.CompositionModelDto{
		Elements:          elements,
		ModuleCount:       dto.ModuleCount,
		ReuseParam:        composition.ReuseFromPercent(dto.ReuseParam),
		ProductivityLevel: getProductivityLevel(dto.ProductivityLevel),
		AveragePay:        dto.AveragePay,
		TeamConfig:        convertTeamConfig(dto.TeamConfig),
	}
	return a.compositionService.CalculateCompositionModel(serviceDto)
}

// CalculateArchitectureModel calculates architecture model
func (a *App) CalculateArchitectureModel(dto ArchitectureModelDtoFrontend) architecture.ArchitectureModelResult {
	config := architecture.ArchitectureAttributesConfig{
		PERS: getAttributeLevel(dto.ArchitectureAttributes["PERS"]),
		RCPX: getAttributeLevel(dto.ArchitectureAttributes["RCPX"]),
		RUSE: getAttributeLevel(dto.ArchitectureAttributes["RUSE"]),
		PDIF: getAttributeLevel(dto.ArchitectureAttributes["PDIF"]),
		PREX: getAttributeLevel(dto.ArchitectureAttributes["PREX"]),
		FCIL: getAttributeLevel(dto.ArchitectureAttributes["FCIL"]),
		SCED: getAttributeLevel(dto.ArchitectureAttributes["SCED"]),
	}
	serviceDto := architecture.ArchitectureModelDto{
		ArchitectureAttributes: &config,
		TeamConfig:             convertTeamConfig(dto.TeamConfig),
		Kloc:                   dto.Kloc,
		AveragePay:             dto.AveragePay,
	}
	return a.architectureService.CalculateArchitectureModel(serviceDto)
}
