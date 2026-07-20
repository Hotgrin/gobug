export namespace main {
	
	export class RunResult {
	    stdout: string;
	    stderr: string;
	    success: boolean;
	    explanation: string;
	
	    static createFrom(source: any = {}) {
	        return new RunResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.stdout = source["stdout"];
	        this.stderr = source["stderr"];
	        this.success = source["success"];
	        this.explanation = source["explanation"];
	    }
	}

}

