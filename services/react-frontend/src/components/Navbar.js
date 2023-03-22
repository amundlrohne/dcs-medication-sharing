import React, { useEffect, useState } from "react";

import "../css/Navbar.css";

import { resolveURL } from "../util/resolveURL";

const Navbar = () => {
    const [logoutButton, showLogoutButton] = useState(false);

    useEffect(async () => {
        let pathname = window.location.pathname;
        if (pathname !== "/login") {
            let cookieData = await getCookie();
            console.log(cookieData);

            if (cookieData.message === "success") {
                showLogoutButton(true);
            }
        }
    }, []);

    const logoutUser = async () => {
        let response = await deleteCookie();

        window.location = "/login";
    };

    const deleteCookie = async () => {
        let url = `${resolveURL("healthcare-provider")}/healthcare-provider/`;
        const response = await fetch(url, {
            method: "DELETE",
            mode: "cors",
            credentials: "include",
        });

        return response.json();
    };

    const getCookie = async () => {
        let url = `${resolveURL(
            "healthcare-provider"
        )}/healthcare-provider/current`;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include",
        });

        return response.json();
    };

    return (
        <div className="navbar">
            <p>MedShare</p>
            {logoutButton && (
                <button
                    onClick={() => {
                        logoutUser();
                    }}
                    className="logout-button"
                >
                    Logout
                </button>
            )}
        </div>
    );
};

export default Navbar;
