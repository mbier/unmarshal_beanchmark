# Unmarshal Benchmark 

| Test                                        | Count   |             |            |               |
|---------------------------------------------|--------:|------------:|-----------:|--------------:|
| Benchmark_unmarshalAvroHamba-12             | 5714372 | 216.5 ns/op | 48 B/op    | 1 allocs/op   |
| Benchmark_unmarshalProto-12                 | 3673202 | 322.9 ns/op | 104 B/op   | 3 allocs/op   |
| Benchmark_unmarshalAvroLinkedin-12          | 2133409 | 558.8 ns/op | 544 B/op   | 10 allocs/op  |
| Benchmark_unmarshalJson-12                  | 918817  | 1275 ns/op  | 264 B/op   | 6 allocs/op   |
| Benchmark_unmarshalProto_dynamicOfficial-12 | 861507  | 1478 ns/op  | 640 B/op   | 12 allocs/op  |
| Benchmark_unmarshalJsonDynamic-12           | 701875  | 1702 ns/op  | 784 B/op   | 23 allocs/op  |
| Benchmark_unmarshalProto_dynamicJhump-12    | 630722  | 1956 ns/op  | 608 B/op   | 22 allocs/op  |
| Benchmark_unmarshalAvroGenerated-12         | 67849   | 17984 ns/op | 14208 B/op | 252 allocs/op |
