import React from "react";
import { GetTaskData } from "../../action/Tasks/get_by_id";
import Navbar from "../Components/Navbar";
import CodeEditor from '@uiw/react-textarea-code-editor';
import { GetTests } from "../../action/Tasks/get_tests";
import { IsItSolved } from "../../action/Tasks/is_solved";
import { FaCheckCircle } from 'react-icons/fa'; // Assuming you are using Font Awesome icons

function Task() {
    const [id, setId] = React.useState(null);
    const [data, setData] = React.useState(null);
    const [code, setCode] = React.useState(
        ``
      );
    const [test, setTest] = React.useState(null)
    const [added, setAdded] = React.useState(false)
    const [output, setOutput] = React.useState(null)
    const [result, setResult] = React.useState(null)
    const [solved, setSolved] = React.useState(false)
    

    React.useEffect(() => {
        const currentRoute = window.location.pathname;

        if (currentRoute.startsWith("/t/")) {
            const remainingPart = currentRoute.substring(3); 
            console.log("Remaining part:", remainingPart);
            setId(remainingPart);
        } else {
            console.log("Current route doesn't start with '/t/'");
        }
                
    }, []);

    React.useEffect(() => {
        if (test !== null && !added){
        let text = `\n#Input code below. Example of test: 
        #Input: ${test.input} 
        #Expected output: ${test.output}`;
        setCode(prev => prev + text);
        setAdded(true)
    } 
    }, [test]);
    

    React.useEffect(() => {
        async function fetchTaskData() {
            try {
                const resp = await GetTaskData(id);
                console.log(resp)
                if (resp.message !== "sql: no rows in result set"){
                    setData(resp);
                }
                const respTest = await GetTests(id)
                console.log(respTest)
                setTest(respTest)

                const idDid = await IsItSolved(localStorage.getItem("login"),id)
                setSolved(idDid)

            } catch (error) {
                console.error("Error fetching task data:", error);
            }
        }
        if (id) {
            fetchTaskData();
        }
    }, [id]);

    async function runCode() {
        try {
            const response = await fetch('http://localhost:8080/compile/python', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ code })
            });
            
            const responseData = await response.json();
            console.log(responseData);
            setResult(responseData)
        } catch (error) {
            console.error('Error running code:', error);
        }
    }

    async function submitCode() {
        try {
            const response = await fetch(`http://localhost:8080/task/${id}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem("token")
                },
                body: JSON.stringify({ code })
            });
            
            const responseData = await response.json();
            if (responseData.message == "success" && responseData.total === responseData.passed){
                setSolved(true)
            }
            setResult(responseData)
        } catch (error) {
            console.error('Error running code:', error);
        }
    }

    return (
        <div className="h-screen">
            <Navbar />
            <div className="bg-base-200 h-5/6 flex">
                <div className="w-1/2">
                    {data && (
                        <div className="border rounded-lg shadow-lg p-4">
                             <div style={{ display: 'flex', alignItems: 'center' }}>
                                <h1 className="text-2xl font-bold" style={{ display: 'inline-block', margin: 0 }}>
                                    {data.id}. {data.title} 
                                </h1>                           
                                {solved && <FaCheckCircle className="ml-2" style={{ color: 'green', marginLeft: '5px' }} />}
                            </div>

                            <p className="text-gray-600">by {data.author}</p>
                            <p className="mt-4">{data.description}</p>
                            <img src={data.image} alt="Task Image" className="mt-4 rounded-lg" />
                            <div className="mt-4">
                                <p className="font-bold">Examples:</p>
                                
                                {data!==null && data.example && data.example.map((exampleArray, index) => (
                                <div key={index} className="mt-4">
                                    {exampleArray.split(";").map((example, subIndex) => (
                                        <p key={subIndex} className="mt-2">{example}</p>
                                    ))}
                                </div>
                            ))}
                            </div>
                        </div>
                    )}
                </div>
                <div className="w-1/2">
                    <div className="border rounded-lg shadow-lg p-1 max-h-2/3 overflow-y-scroll">
                        <CodeEditor
                        className="max-h-2/3"
                            value={code}
                            language="python"
                            placeholder="Please enter Python code."
                            onChange={(evn) => setCode(evn.target.value)}
                            minHeight={500}
                            maxHeight={500}
                            style={{
                                backgroundColor: "#161b22",
                                fontFamily: 'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace',
                            }}
                            />
                   
                    </div>
                    <div className="p-1">
                            <div className="flex justify-center align-center mt-5">
                                {localStorage.getItem("role") !== "teacher" &&
                                <button className="btn btn-active btn-primary" onClick={submitCode}>Submit code</button>    }    
                                <button className="btn btn-active btn-accent ml-5" onClick={runCode}>Run code</button>                                                
                            </div>
                            {output !== null && result === null &&
                                <div className="flex justify-center align-center mt-5">
                                    <div className="mockup-code w-1/2">
                                        <pre data-prefix="$"><code>{output}</code></pre>
                                    </div>
                                </div>
                            }
                            {result !== null && result.message !== "success" &&
                                <div className="flex justify-center align-center mt-5">
                                    <div className="mockup-code w-1/2">
                                        <pre data-prefix="$"><code>{result.message}</code></pre>
                                    </div>
                                </div>
                            }
                            {result !== null && result.message === "success" &&
                            <div style={{ display: "flex", alignItems: "center" }} className="flex justify-center align-center mt-5">
                            <span>{result.passed} passed</span>
                            <progress
                              className="progress progress-success"
                              value={result.passed}
                              max={result.total}
                              style={{ margin: "0 5px", width: "150px" }}
                            ></progress>
                            <span>{result.total} total</span>
                          </div>
                          
                            }

                        
                    </div>
                    
                </div>
            </div>
        </div>
    );
}

export default Task;
