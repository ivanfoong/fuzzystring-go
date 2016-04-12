package fuzzystring_test

import (
  "testing"
  "github.com/ivanfoong/fuzzystring-go/fuzzystring"
  "fmt"
)

func TestCompareString(t *testing.T) {
  var testCases = []struct {
    inputs []string
    output float64
  }{
    {[]string{"Healed", "Sealed"}, 0.64},
    {[]string{"Healed", "Healthy"}, 0.44},
    {[]string{"Healed", "Heard"}, 0.36},
    {[]string{"Healed", "Herded"}, 0.32},
    {[]string{"Healed", "Help"}, 0.2},
    {[]string{"Healed", "Sold"}, 0.0},
    {[]string{"Web Database Applications with PHP & MySQL", "Web Database Applications"}, 0.79},
    {[]string{"Web Database Applications with PHP & MySQL", "PHP Web Applications"}, 0.68},
    {[]string{"Web Database Applications with PHP & MySQL", "Web Aplications"}, 0.52},
    {[]string{"Creating Database Web Applications with PHP and ASP", "Web Database Applications"}, 0.68},
    {[]string{"Creating Database Web Applications with PHP and ASP", "PHP Web Applications"}, 0.58},
    {[]string{"Creating Database Web Applications with PHP and ASP", "Web Aplications"}, 0.44},
    {[]string{"Building Database Applications on the Web Using PHP3", "Web Database Applications"}, 0.67},
    {[]string{"Building Database Applications on the Web Using PHP3", "PHP Web Applications"}, 0.53},
    {[]string{"Building Database Applications on the Web Using PHP3", "Web Aplications"}, 0.43},
    {[]string{"Building Web Database Applications with Visual Studio 6", "Web Database Applications"}, 0.65},
    {[]string{"Building Web Database Applications with Visual Studio 6", "PHP Web Applications"}, 0.46},
    {[]string{"Building Web Database Applications with Visual Studio 6", "Web Aplications"}, 0.41},
    {[]string{"Web Application Development With PHP", "Web Database Applications"}, 0.46},
    {[]string{"Web Application Development With PHP", "PHP Web Applications"}, 0.63},
    {[]string{"Web Application Development With PHP", "Web Aplications"}, 0.51},
    {[]string{"WebRAD: Building Database Applications on the Web with Visual FoxPro and Web Connection", "Web Database Applications"}, 0.47},
    {[]string{"WebRAD: Building Database Applications on the Web with Visual FoxPro and Web Connection", "PHP Web Applications"}, 0.32},
    {[]string{"WebRAD: Building Database Applications on the Web with Visual FoxPro and Web Connection", "Web Aplications"}, 0.29},
    {[]string{"Structural Assessment: The Role of Large and Full-Scale Testing", "Web Database Applications"}, 0.1},
    {[]string{"Structural Assessment: The Role of Large and Full-Scale Testing", "PHP Web Applications"}, 0.06},
    {[]string{"Structural Assessment: The Role of Large and Full-Scale Testing", "Web Aplications"}, 0.06},
    {[]string{"How to Find a Scholarship Online", "Web Database Applications"}, 0.08},
    {[]string{"How to Find a Scholarship Online", "PHP Web Applications"}, 0.09},
    {[]string{"How to Find a Scholarship Online", "Web Aplications"}, 0.1},
    {[]string{"Plain Crackers", "Plain Crackers W Oats  J' Pk"}, 0.81},
    {[]string{"Plain Crackers", "Plain Crackers - J'Pack"}, 0.86},
  }
  for _, tt := range testCases {
    text1 := tt.inputs[0]
    text2 := tt.inputs[1]
    expectedScore := tt.output
    score := fuzzystring.CompareString(text1, text2)
    if fmt.Sprintf("%.2f", expectedScore) != fmt.Sprintf("%.2f", score) {
      t.Errorf("expected CompareString(%s, %s) to equal `%.2f`, got `%.2f` instead", text1, text2, expectedScore, score)
    }
  }
}