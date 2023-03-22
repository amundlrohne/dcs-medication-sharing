import React, {useEffect, useState} from 'react'


import '../css/ConsentForm.css'


const ConsentForm = () => {

    const [senderHP, setSenderHP] = useState(['sHP1', 'sHP2'])
    const [receiverHP, setReceiverHP] = useState(['RHP1', 'sHP1'])
    const [selectedSenderHP, setSelectedSenderHP] = useState([])
    const [selectedReceiverHP, setSelectedReceiverHP] = useState([])
    const [validationMsg, setValidationMsg] = useState([])

    function handleSubmit(e){
        let senderHP = e.target.senderHP.value;
        let receiverHP = e.target.receiverHP.value;
        if (senderHP === receiverHP){
            e.preventDefault();
            setValidationMsg(["Sender and receiver Healthcare Providers must be different"]);
            return;
        }
        
        e.preventDefault();
        createConcent();
        
    }

 


    // get the id
    async function getHPByName(name){
        let fetchUrl = 'http://172.29.73.112:8280/health-provider/name/'+name;
        const res = await fetch(fetchUrl);
        const d = await res.json();
        return d.data.data;
    }

    async function postConsent(senderPK, receiverPK){
        
        let postUrl  = 'http://172.29.73.112:8180/consent';

        let d = new Object();

        d.ToPublicKey = senderPK;
        d.FromPublicKey = receiverPK;
        d.ExpDate = dateNowPlusOneMonth();
        d.DateCreated = dateNow();

        let jsonBody = JSON.stringify(d)
        console.log(jsonBody);
        const req = await fetch(postUrl,{
            method: 'POST',
            headers: {
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Allow-Methods": "*",
                "Access-Control-Request-Headers": "*",
                "Access-Control-Request-Method": "*",
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonBody) 
        })

        console.log(req.status);
        console.log(req.json())
        
    }

    function dateNow(){

        let today = new Date();
        let dd = today.getDate();
        let mm = today.getMonth()+1; 
        let yyyy = today.getFullYear();
        if(dd<10) 
        {
            dd='0'+dd;
        } 

        if(mm<10) 
        {
            mm='0'+mm;
        } 

        today = yyyy+'-'+mm+'-'+dd
        return today;
    }

    function dateNowPlusOneMonth(){
        let now = new Date();
        let current = new Date(now.getFullYear(), now.getMonth()+1, 1);

        let dd = current.getDate();
        let mm = current.getMonth()+1; 
        let yyyy = current.getFullYear();
        if(dd<10) 
        {
            dd='0'+dd;
        } 

        if(mm<10) 
        {
            mm='0'+mm;
        } 
        let todayPlusOneMonth = yyyy+'-'+mm+'-'+dd
        return todayPlusOneMonth;
    }

           
    useEffect(()=>{
        async function fetchHP(){
            //Change url
            let fetchUrl = 'http://172.29.73.112:8280/health-provider/all' ;
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
            setSelectedSenderHP(d[0].name);
            setSelectedReceiverHP(d[0].name);
            
        };
        fetchHP();
    }, [])


    function handleChangeSenderHP(e){
        setSelectedSenderHP([e.target.value])
        
    }

    function handleChangeReceiverHP(e){
        setSelectedReceiverHP([e.target.value]);
       
    }

    async function createConcent(){
        console.log("Value sender: " + selectedSenderHP);
        console.log("Value receiver: " + selectedReceiverHP);

        let senderId = await getHPByName(selectedSenderHP);
        let receiverId = await getHPByName(selectedReceiverHP);
        
        const res = await postConsent(senderId, receiverId);
        
    }

    return (        

    <div className='ConsentForm'>
        <form className='consent-form' onSubmit={handleSubmit} action='#'>
            <label htmlFor="senderHP">Sender Healthcare Provider</label>
            <select name='senderHP' onChange={(e) => { handleChangeSenderHP(e) }}>
                {senderHP.map((value, index) => <option key={index} name={value}>{value}</option>)}
            </select>
            <br></br>
            <label htmlFor="receiverHP">Reciever Healthcare Provider</label>
            <select name='receiverHP' onChange={(e) => { handleChangeReceiverHP(e) }}>
                {receiverHP.map((value, index) => <option key={index} name={value}>{value}</option>)}
            </select>
            <br></br>
            <label htmlFor="agreeConcent"> I consent to share my medication record between the healthcare providers
            <input type="checkbox" name="agreeConcent" value="agree" required/>
            </label><br></br>

            
            {/* Validation error message  */}
            {validationMsg.map((val, idx) => <div key={idx}><h3>{val}</h3></div>)}
            
            <input type="submit" value="Create Consent" />
        </form>
    </div>
    
    )

}



export default ConsentForm