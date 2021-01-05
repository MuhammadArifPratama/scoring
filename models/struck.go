package models

type PMK struct {
	ProspectID		 	 string			`json:"prospect_id,omitempty"`
	BirthDate			 string			`json:"birth_date,omitempty"`
	MaritalStatus 		 string			`json:"marital_status,omitempty"`
	MonthlyFixedIncome 	 int			`json:"monthly_fixed_income,omitempty"`
	WorkSinceMonth		 *int			`json:"worksince_month,omitempty"`
	WorkSinceYear		 int			`json:"worksince_year,omitempty"`
	ProfessionID		 string			`json:"profession_id,omitempty"`
	Tenor			     int			`json:"tenor,omitempty"`
}
