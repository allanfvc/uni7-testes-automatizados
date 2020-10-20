Feature: calculadora

  Scenario: Sum 30
    Given numbers 1.0 and 1.0
    Then sum should be 2.0
  
  Scenario: Sum 0.24
    Given numbers 0.2 and 0.04
    Then sum should be 0.24