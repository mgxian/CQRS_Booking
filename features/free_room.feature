Feature: free room
  In order to book a room
  As a consumer
  I need to know free rooms

  Scenario: have free rooms when book one day
    Given the shanghai room was booked as the following:
      | arrival    | departure  |
      | 2020-10-19 | 2020-10-20 |
      | 2020-10-21 | 2020-10-22 |
    When get free rooms arrival at 2020-10-20 and departure at 2020-10-21
    Then the free rooms should contains the shanghai room

  Scenario: have no free rooms when book one day
    Given the shanghai room was booked as the following:
      | arrival    | departure  |
      | 2020-10-19 | 2020-10-20 |
      | 2020-10-20 | 2020-10-21 |
    When get free rooms arrival at 2020-10-20 and departure at 2020-10-21
    Then the free rooms should not contains the shanghai room

  Scenario: have free rooms when book more than one day
    Given the shanghai room was booked as the following:
      | arrival    | departure  |
      | 2020-10-19 | 2020-10-22 |
      | 2020-10-26 | 2020-10-29 |
    When get free rooms arrival at 2020-10-22 and departure at 2020-10-26
    Then the free rooms should contains the shanghai room

  Scenario: have no free rooms when book more than one day
    Given the shanghai room was booked as the following:
      | arrival    | departure  |
      | 2020-10-19 | 2020-10-25 |
      | 2020-10-26 | 2020-10-29 |
    When get free rooms arrival at 2020-10-25 and departure at 2020-10-27
    Then the free rooms should not contains the shanghai room
