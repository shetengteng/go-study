module demo

go 1.14

require stt v0.0.0

// 在 go.mod 中使用 replace 将 stt 指向 ./stt
replace stt => ./stt