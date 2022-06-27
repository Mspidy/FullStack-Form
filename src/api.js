function Api(){
    function submitForm(){
        console.log("Submit All the form data")
    }
    return(
        <div>
            <form>
                <input type="text" placeholder="Enter First Name"/>
                <br/>
                <br/>
                <input type="text" placeholder="Enter Last Name"/>
                <br/>
                <br/>
                <button onSubmit={submitForm}>Submit</button>
            </form>
        </div>
    )
}
export default Api;