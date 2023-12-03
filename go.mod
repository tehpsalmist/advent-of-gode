module 8thday.dev/aog

go 1.21.4

require 8thday.dev/aog/exercises v0.0.0

require 8thday.dev/aog/utils v0.0.0 // indirect

replace (
	8thday.dev/aog/exercises => ./exercises
	8thday.dev/aog/utils => ./utils
)
