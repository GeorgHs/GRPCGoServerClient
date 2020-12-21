protoc --go_out=plugins=grpc:. greet/greetpb/greet.proto
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.