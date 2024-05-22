function Search(props){
    return (
        <div className="join">
            <div>
                <div>
                <input className="input input-bordered join-item" placeholder="Search"/>
                </div>
            </div>
            <select className="select select-bordered join-item">
                <option disabled selected>Difficulty</option>
                <option>Easy</option>
                <option>Medium</option>
                <option>Hard</option>
            </select>
            <div className="indicator">
                <button className="btn join-item">Search</button>
            </div>
        </div>
    )
}

export default Search 