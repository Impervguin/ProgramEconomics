package chars

type BudgetCharacteristics struct {
	Analysis                float64
	Design                  float64
	Programming             float64
	TestPlanning            float64
	Verification            float64
	Office                  float64
	ConfigurationManagement float64
	Documentation           float64
}

var (
	BudgetCharacteristicsCoefficients BudgetCharacteristics = BudgetCharacteristics{
		Analysis:                0.04,
		Design:                  0.12,
		Programming:             0.44,
		TestPlanning:            0.06,
		Verification:            0.14,
		Office:                  0.07,
		ConfigurationManagement: 0.07,
		Documentation:           0.06,
	}
)

// budget in money
func CalculateBudgetCharacteristics(work, budgetPerWork float64) BudgetCharacteristics {
	totalBudget := work * budgetPerWork
	return BudgetCharacteristics{
		Analysis:                totalBudget * BudgetCharacteristicsCoefficients.Analysis,
		Design:                  totalBudget * BudgetCharacteristicsCoefficients.Design,
		Programming:             totalBudget * BudgetCharacteristicsCoefficients.Programming,
		TestPlanning:            totalBudget * BudgetCharacteristicsCoefficients.TestPlanning,
		Verification:            totalBudget * BudgetCharacteristicsCoefficients.Verification,
		Office:                  totalBudget * BudgetCharacteristicsCoefficients.Office,
		ConfigurationManagement: totalBudget * BudgetCharacteristicsCoefficients.ConfigurationManagement,
		Documentation:           totalBudget * BudgetCharacteristicsCoefficients.Documentation,
	}
}

// work in human-months
func CalculateBudgetWorkCharacteristics(work float64) BudgetCharacteristics {
	return BudgetCharacteristics{
		Analysis:                work * BudgetCharacteristicsCoefficients.Analysis,
		Design:                  work * BudgetCharacteristicsCoefficients.Design,
		Programming:             work * BudgetCharacteristicsCoefficients.Programming,
		TestPlanning:            work * BudgetCharacteristicsCoefficients.TestPlanning,
		Verification:            work * BudgetCharacteristicsCoefficients.Verification,
		Office:                  work * BudgetCharacteristicsCoefficients.Office,
		ConfigurationManagement: work * BudgetCharacteristicsCoefficients.ConfigurationManagement,
		Documentation:           work * BudgetCharacteristicsCoefficients.Documentation,
	}
}
