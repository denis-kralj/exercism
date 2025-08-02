module Leap
open System

let leapYear (year: int): bool = DateTime.IsLeapYear(year)