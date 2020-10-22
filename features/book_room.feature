Feature: book room
  In order to get a room
  As a consumer
  I want book a room

  Scenario: book room
    Given the shanghai room
    And the shanghai room is free arrival at 2020-10-20 and departure at 2020-10-21
    When will book the shanghai room arrival at 2020-10-20 and departure at 2020-10-21
    Then the shanghai room is not free arrival at 2020-10-20 and departure at 2020-10-21
