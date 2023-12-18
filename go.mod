module github.com/MuCephei/AdventOfCode2023

go 1.21.5

require (
	internal/five v0.0.0-00010101000000-000000000000
	internal/four v0.0.0-00010101000000-000000000000
	internal/general v1.0.0
	internal/one v0.0.0-00010101000000-000000000000
	internal/three v0.0.0-00010101000000-000000000000
	internal/two v0.0.0-00010101000000-000000000000
)

replace internal/general => ./internal/general

replace internal/one => ./internal/one

replace internal/two => ./internal/two

replace internal/three => ./internal/three

replace internal/four => ./internal/four

replace internal/five => ./internal/five
