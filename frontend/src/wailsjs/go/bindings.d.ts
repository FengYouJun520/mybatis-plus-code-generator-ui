export interface go {
  "codegen": {
    "Manager": {
		CodeGenerate(arg1:ConfigContext):Promise<string|Error>
    },
  }

  "main": {
    "App": {
		Execute(arg1:ConfigContext):Promise<string|Error>
		GetOs():Promise<string>
		Greet(arg1:string):Promise<string>
		PingDb(arg1:DataSourceConfig):Promise<Array<DatabaseOptions>|Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
