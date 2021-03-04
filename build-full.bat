call npm run build --prefix ./pkg/server/ui 
go generate ./cmd && go build -o ./dist/d2-item-sorter.exe ./cmd