module github.com/MuCephei/AdventOfCode2023

go 1.21.5

require internal/general v1.0.0

replace internal/general => ./internal/general

require (
	internal/one v1.0.0
	internal/two v0.0.0-00010101000000-000000000000
)

replace internal/one => ./internal/one

replace internal/two => ./internal/two

replace internal/three => ./internal/three
