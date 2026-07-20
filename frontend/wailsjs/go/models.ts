export namespace config {
	
	export class Config {
	    apiKey: string;
	    model: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.apiKey = source["apiKey"];
	        this.model = source["model"];
	    }
	}

}

export namespace main {
	
	export class RunResult {
	    stdout: string;
	    stderr: string;
	    success: boolean;
	    explanation: string;
	    explainedByAI: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RunResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.stdout = source["stdout"];
	        this.stderr = source["stderr"];
	        this.success = source["success"];
	        this.explanation = source["explanation"];
	        this.explainedByAI = source["explainedByAI"];
	    }
	}

}

