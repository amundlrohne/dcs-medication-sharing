import React, {useEffect, useState} from 'react'


import '../css/ConsentForm.css'


const ConsentForm = () => {

    const [senderHP, setSenderHP] = useState(['sHP1', 'sHP2'])
    const [receiverHP, setReceiverHP] = useState(['RHP1', 'sHP1'])
    const [validationMsg, setValidationMsg] = useState([])

    function handleSubmit(e){
        let senderHP = e.target.senderHP.value;
        let receiverHP = e.target.receiverHP.value;
        if (senderHP==receiverHP){
            e.preventDefault();
            setValidationMsg(["Sender and receiver Healthcare Providers must be different"]);
            return;
        }
        //Submit the data
            
    }

    async function getHPId(){
        let senderHP = e.target.senderHP.value;
        
    }

    async function fetchHP(){
        //Change url
        let fetchUrl = 'http://localhost:8280/health-provider/all' ;
        let healthCareProviderNames = [] ;
        let d = [];

        const res = await fetch(fetchUrl);
        d = await res.json();
        d = d.data.data;
        for (let i =0 ; i < d.length ; i++){
            healthCareProviderNames.push(d[i].name);
        }
        setSenderHP(healthCareProviderNames);
        setReceiverHP(healthCareProviderNames);
        
    }

        
    useEffect(async ()=>{
        await fetchHP();
    }, [])


    return (        

    <div className='ConsentForm'>
        <form className='consent-form' onSubmit={handleSubmit} action='/test' method='POST' >
            <label for="senderHP">Sender Healthcare Provider</label>
            <select name='senderHP'>
                {senderHP.map((value, index) => <option key={index} name={value}>{value}</option>)}
            </select>
            <br></br>
            <label for="receiverHP">Reciever Healthcare Provider</label>
            <select name='receiverHP'>
                {receiverHP.map((value, index) => <option key={index} name={value}>{value}</option>)}
            </select>
            <br></br>
            <label for="agreeConcent"> I consent to share my medication record between the healthcare providers
            <input type="checkbox" name="agreeConcent" value="agree" required/>
            </label><br></br>

            {/* Validation error message  */}
            {validationMsg.map((val, idx) => <div key={idx}><h3>{val}</h3></div>)}
            
            <input type="submit" value="Create Consent"/>
        </form>
    </div>
    
    )

}



export default ConsentForm