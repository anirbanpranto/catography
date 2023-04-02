import axios from "axios"
import React, { useState, useEffect, useRef } from "react"
import { MdUpload } from "react-icons/md";
import mapboxgl from 'mapbox-gl'; // eslint-disable-line import/no-webpack-loader-syntax
import { Center, Box, Button, VStack } from "@chakra-ui/react";
import 'mapbox-gl/dist/mapbox-gl.css';
import SimpleSidebar from "./Sidebar";
import { v4 } from "uuid"


interface image_file {
    Url: string
    Time: string
    Unsigned: string
    Lon: number
    Lat: number
}

const Map = () => {
    mapboxgl.accessToken = import.meta.env.VITE_MAPBOX_GL_TOKEN
    const mapContainer = useRef(null);
    const map = useRef(null);
    const [lng, setLng] = useState(-77.432);
    const [lat, setLat] = useState(25.0306);
    const [zoom, setZoom] = useState(9);

    useEffect(() => {
        if (map.current) return; // initialize map only once
        //@ts-ignore
        map.current = new mapboxgl.Map({
            //@ts-ignore
            container: mapContainer.current,
            style: 'mapbox://styles/mapbox/streets-v12',
            center: [lng, lat],
            zoom: zoom
        })
        //@ts-ignore
        map.current?.addControl(new mapboxgl.GeolocateControl({
            positionOptions: {
                enableHighAccuracy: true
            },
            // When active the map will receive updates to the device's location as it changes.
            trackUserLocation: true,
            // Draw an arrow next to the location dot to indicate which direction the device is heading.
            showUserHeading: true
        }));

        var socket_connection = new WebSocket("ws://localhost:8080/v1/images/connect");

        socket_connection.addEventListener("open", (event) => {
            socket_connection.send("Hello Server!");
        });

        // Listen for messages
        socket_connection.addEventListener("message", async (event) => {
            console.log(event.data)
            await getData();
        });

        const getData = async () => {
            await axios.get('http://localhost:8080/v1/images/').then(async (res) => {
                setImages(res.data);
                if(res.data){
                    await updateMap(res.data);
                }
            });
            //@ts-ignore
            // Load an image from an external URL.
        };

        getData();
        //@ts-ignore
    }, []);

    const updateMap = async (images: image_file[]) => {
        for (let i = 0; i < images.length; i++) {
            //@ts-ignore
            map.current.loadImage(
                images[i].Url,
                //@ts-ignore
                (error, image) => {
                    if (error) {
                        console.log("Error")
                    }

                    // Add the image to the map style.
                    //@ts-ignore
                    map.current.addImage(images[i].Unsigned, image);

                    //@ts-ignore
                    map.current.addSource(images[i].Unsigned, {
                        'type': 'geojson',
                        'data': {
                            'type': 'FeatureCollection',
                            'features': [
                                {
                                    'type': 'Feature',
                                    'geometry': {
                                        'type': 'Point',
                                        'coordinates': [images[i].Lon, images[i].Lat]
                                    }
                                }
                            ]
                        }
                    });
                    //@ts-ignore
                    // Add a layer to use the image to represent the data.
                    map.current.addLayer({
                        'id': i.toString(),
                        'type': 'symbol',
                        'source': images[i].Unsigned, // reference the data source
                        'layout': {
                            'icon-image': images[i].Unsigned, // reference the image
                            'icon-size': 0.25
                        }
                    });
                }
            )
        }
    }

    const renderImages = () => {
        const image_arr = images.map((image, idx) => {
            return <img key={idx} src={image.Url} />
        });
        return image_arr;
    }

    const [images, setImages] = useState<image_file[]>([]);
    return (
        <>
            <div className="map-container" ref={mapContainer}>
            </div>
            <SimpleSidebar children={[]} />
        </>
    )
}

export default Map;