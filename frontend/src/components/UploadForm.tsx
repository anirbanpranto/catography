import React from "react";
import { useState, useEffect } from "react"
import axios from "axios";
import { Center, FormControl, FormLabel, Input, FormHelperText, Box, Button } from "@chakra-ui/react";
import SimpleSidebar from "./Sidebar";

const UploadForm = () => {
    const uploadFile = (file: FileList | null) => {
        //@ts-ignore
        setFile(file[0]);
    };

    const [remarks, setRemarks] = useState("");
    const [file, setFile] = useState(null);

    return (
        <>
            <SimpleSidebar children={[]} />
            <Center>
                <FormControl style={{width:"inherit", marginTop:"300px", marginLeft:"20px", marginRight:"20px"}}>
                    <FormLabel>Remarks</FormLabel>
                    <Input value={remarks} onChange={(e) => { setRemarks(e.target.value) }} type='text' />
                    <FormHelperText>Please write something.</FormHelperText>
                    <FormLabel>Image</FormLabel>
                    <Input className="InputButton" type="file" onChange={(e) => uploadFile(e.target.files)} />
                    <Button onClick={() => {
                        const formData = new FormData();
                        if (file) {
                            navigator.geolocation.getCurrentPosition((position) => {
                                formData.append('file', file);
                                formData.append('remarks', remarks);
                                formData.append('lat', position.coords.latitude.toString());
                                formData.append('lon', position.coords.longitude.toString());
                                axios.post('http://localhost:8080/v1/images/', formData, {
                                    headers: {
                                        'Content-Type': 'multipart/form-data',
                                    },
                                });
                            })
                        }
                    }}
                    style={{marginTop:"20px"}}
                    >Submit</Button>
                </FormControl>
            </Center>
        </>
    )
}

export default UploadForm;