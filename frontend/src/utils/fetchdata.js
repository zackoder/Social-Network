export async function fetchdata(obj) {
    const response = await fetch(obj.URl,{ 
        method: obj.method,
        credentials:"include",
        body:obj.body,
    })
    const data = await response.json()
    return [data,response]
}