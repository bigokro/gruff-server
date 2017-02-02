package gruff

/*
 * A Claim is a proposed statement of fact
 *
 * According to David Zarefsky (https://www.thegreatcoursesplus.com/argumentation/argument-analysis-and-diagramming) there are 4 types:
 * - Fact: Al Gore received more popular votes than George Bush in the 2000 election
 * - Definition: Capital execution is murder
 * - Value: Environmental protection is more important than economic growth
 * - Policy: Congress should pass the president's budget
 *
 * Also according to the professor, there are 4 parts to a claim/argument:
 * - Claim
 * - Evidence
 * - Inference
 * - Warrant
 *
 * In loose terms, a Claim here represents his Claim, and Evidence
 * An Argument of type 1 or 2 (truth) is an Inference
 * An Argument of type 3, 4, 5 or 6 is a Warrant
 *
 * Complex Claims:
 * - Series: Because of X, Y happened, which caused Z --> Not modeled in Gruff
 * - Convergent: Airline travel is becoming more unpleasant because of X, Y, Z, P, D, and Q --> Supported by standard Gruff structure
 * - Parallel: Same as convergent, except that any one argument is enough --> Supported by standard Gruff structure
 */
type Claim struct {
	Identifier
	Title       string     `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description string     `json:"desc" valid:"length(3|4000)"`
	Truth       float64    `json:"truth"`
	ProTruth    []Argument `json:"protruth,omitempty"`
	ConTruth    []Argument `json:"contruth,omitempty"`
	Links       []Link     `json:"links,omitempty"`
	Contexts    []Context  `json:"contexts,omitempty"  gorm:"many2many:claim_contexts;"`
	Values      []Value    `json:"values,omitempty"  gorm:"many2many:claim_values;"`
	Tags        []Tag      `json:"tags,omitempty"  gorm:"many2many:claim_tags;"`
}
