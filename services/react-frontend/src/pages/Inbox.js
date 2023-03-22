import React, {useEffect, useState} from 'react'

import '../css/Inbox.css'

import { resolveURL } from "../util/resolveURL";

const Inbox = () => {
    const [consentRequests, setConsentRequests] = useState([])
    const [consentIncoming, setConsentIncoming] = useState([])
    
    const [requestsVisibility, showRequests] = useState(true)
    const [incomingVisibility, showIncoming] = useState(false)

    const [requestsSelected, setRequestsSelected] = useState(['#5CDB95', '1px solid #5CDB95', 'bold'])
    const [incomingSelected, setIncomingSelected] = useState([])

    useEffect(async () => {

        let cookieData = await getCookie();
        console.log(cookieData);

        if (cookieData.message !== "success") {
            window.location = "/not-authorized";
        } else {
            // Fetch current providers JWT
            let fetchCurrentProvider = await getCookie()
            let token = fetchCurrentProvider.data.data

            // Fetch current providers ID
            let provider = await getProviderByToken(token)
            let provider_id = provider.data.data.ID

            // Fetch consents used as requests
            // These consents will be used for writing medication records
            let fetchConsentRequests = await getConsentRequests(provider_id)
            let dataRequests = fetchConsentRequests.data.data
            console.log(dataRequests)

            let temp = []
            if (dataRequests !== null) {
                for (var i = 0; i < dataRequests.length; i++) {
                    let consentObj = {
                        consentid: dataRequests[i].ID
                    }
        
                    let bundle = await getBundle(consentObj)
                    let parsedJSON = JSON.parse(bundle.data.data)
        
                    console.log(parsedJSON.total)
        
                    if (parsedJSON.total === 0) { temp.push(dataRequests[i]) }
                }

                setConsentRequests(temp)
            } else {
                setConsentRequests([])
            }

            // Fetch consents used as incoming
            // These consents will be used for reading the filled out medicaiton records
            let fetchConsentIncoming = await getConsentIncoming(provider_id)
            // setConsentIncoming(fetchConsentIncoming.data.data)
            console.log(fetchConsentIncoming)
            let dataIncoming = fetchConsentIncoming.data.data

            if (dataIncoming !== null) {
                for (var i = 0; i < dataIncoming.length; i++) {
                    let consentObj = {
                        consentid: dataIncoming[i].ID
                    }
    
                    console.log(dataIncoming[i].ID)
    
                    let bundle = await getBundle(consentObj)
                    let parsedJSON = JSON.parse(bundle.data.data)
    
                    console.log(parsedJSON)
    
                    if (parsedJSON.total !== 0) { consentIncoming.push(dataIncoming[i]) }
                }
            } else {
                setConsentIncoming([])
            }
        }

    }, [])

    const getBundle = async (data) => {
        let url = `${resolveURL("medication-record")}/medication-record`
        const response = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        return response.json()
    }

    const getConsentRequests = async (id) => {
        let url = `${resolveURL("consent")}/consent/from/` + id;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include"
        })
    
        return response.json()
    }

    const getConsentIncoming = async (id) => {
        let url = `${resolveURL("consent")}/consent/to/` + id;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include"
        })
    
        return response.json()
    }

    const getProviderByToken = async (token) => {
        let url = `${resolveURL("healthcare-provider")}/healthcare-provider/getID/` + token;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include"
        })
    
        return response.json()
    }

    const getCookie = async () => {
        let url = `${resolveURL("healthcare-provider")}/healthcare-provider/current`;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include",
        });
    
        return response.json();
    };

    const selectOption = (id) => {
        if (id === 'requests-tag') {
            setRequestsSelected(['#5CDB95', '1px solid #5CDB95', 'bold'])
            setIncomingSelected([])

            //show requests page
            showRequests(true)
            showIncoming(false)
        }

        if (id === 'incoming-tag') {
            setRequestsSelected([])
            setIncomingSelected(['#5CDB95', '1px solid #5CDB95', 'bold'])

            //show incoming page
            showIncoming(true)
            showRequests(false)
        }
    }

    const viewForm = (id) => {
        window.location.href = '/medication-form?consent_id=' + id
    }

    const viewRecord = (id) => {
        window.location.href = '/view-medication-record?consent_id=' + id
    }

    return(<div>
        <div className='inbox-container'>
            <div className='directory'>
                <p id='directory-title'>Inbox</p>

                <div className='options'>
                    <p id='requests-tag'
                        style={{color: requestsSelected[0], 
                                borderLeft: requestsSelected[1], 
                                fontWeight: requestsSelected[2]}}
                        onClick={() => selectOption("requests-tag")}>Requests</p>
                    <p id='incoming-tag'
                        style={{color: incomingSelected[0], 
                                borderLeft: incomingSelected[1], 
                                fontWeight: incomingSelected[2]}}
                        onClick={() => selectOption("incoming-tag")}>Incoming</p>
                </div>
            </div>

            <div className='options-pages'>
                {requestsVisibility && (
                    <div className='requests-page'>
                        <p id='title'>Requests</p>
                        <div className='consent-list'>
                        {consentRequests.map((consent_elem, idx) => {
                            return (<div className='consent-request-elem' id={consent_elem.ID} key={idx}>
                                        <div className='information-div'>
                                            <p id='information'>{consent_elem.ID}</p>
                                            <p id='information'>{consent_elem.datecreated}</p>
                                        </div>
                                        <button id='view' onClick={() => { viewForm(consent_elem.ID) }}>View</button>
                                    </div>)
                            })}
                        </div>
                    </div>
                )}

{               incomingVisibility && (
                    <div className='incoming-page'>
                        <p id='title'>Incoming</p>
                        <div className='records-list'>
                        {consentIncoming.map((consent_elem, idx) => {
                            return (<div className='consent-request-elem' id={consent_elem.ID} key={idx}>
                                        <div className='information-div'>
                                            <p id='information'>{consent_elem.ID}</p>
                                            <p id='information'>{consent_elem.datecreated}</p>
                                        </div>
                                        <button id='view' onClick={() => { viewRecord(consent_elem.ID) }}>View</button>
                                    </div>)
                            })}
                        </div>
                    </div>
                )}
            </div>
        </div>
    </div>)
}

export default Inbox