# ETH Gas price tracker for Tidbyt
# MIT License
# by Connor Easton, Dec 29 2021

load("render.star", "render")
load("http.star", "http")
load("encoding/base64.star", "base64")
load("cache.star", "cache")

ETHERSCAN_API_KEY = "38MQP8IICBD6UE3A2A7XCNHURTM1KY7AZF"
# key can be obtained for free from https://etherscan.io/myapikey

GAS_PRICE_URL = "https://api.etherscan.io/api?module=gastracker&apikey=" + ETHERSCAN_API_KEY + "&action=gasoracle"
# Also tracks estimated time for a transaction to be confirmed on the blockchain
TIME_TO_VERIFY_URL = "https://api.etherscan.io/api?module=gastracker&action=gasestimate&apikey=" + ETHERSCAN_API_KEY +"&gasprice=2000000000"


# Alternative gas price source (no key needed)
# GAS_PRICE_URL = "https://ethgasstation.info/json/ethgasAPI.json"

GAS_ICON = base64.decode("""
iVBORw0KGgoAAAANSUhEUgAAAAkAAAAJCAYAAADgkQYQAAAAQElEQVQYlWNgYGBguMrA8B8ZM2AD2BRVhnr8x1AEA8gmoSjEZR1eRXv37oVLwtlkKYJJIismHAROTk7/sWFkNQBc0mLyIWghtgAAAABJRU5ErkJggg==
""")
TIME_ICON = base64.decode("""
iVBORw0KGgoAAAANSUhEUgAAAAkAAAAJCAYAAADgkQYQAAAAMklEQVQYlWNgQAL2gVv/wzADOsAmgSGGzMHKxmYCBp+QNVgVYXMvViuwWk2U79Al0CUBE48/cUr118oAAAAASUVORK5CYII=
""")

def main():
    cached_gas = cache.get("gas_price")
    if cached_gas != None:
        print("Hit! Displaying cached data.")
        gas_price = int(cached_gas)

        cached_time = cache.get("time_to_verify")
        time_to_verify = int(cached_time)

        seconds = time_to_verify % (24*3600)
        hours = seconds // 3600
        seconds %= 3600
        minutes = seconds // 60
    else:
        print("Miss! Calling Etherscan API.")
        gas_rep = http.get(GAS_PRICE_URL)
        time_rep = http.get(TIME_TO_VERIFY_URL)
        if (gas_rep.status_code != 200 or time_rep.status_code != 200):
           fail("Etherscan request failed with status %d", gas_rep.status_code)

        if gas_rep.json()["message"] != "OK":
            gas_price = 0
            time_to_verify = 0
            fail("Etherscan API returned error:", gas_rep.json()["result"])

        gas_price = int(gas_rep.json()["result"]["ProposeGasPrice"])# in Gwei
        time_to_verify = int(time_rep.json()["result"]) # in seconds
        
        cache.set("time_to_verify", str(int(time_to_verify)), ttl_seconds=1)
        cache.set("gas_price", str(int(gas_price)), ttl_seconds=1)

        seconds = time_to_verify % (24*3600)
        hours = seconds // 3600
        seconds %= 3600
        minutes = seconds // 60

    return render.Root(
        child = render.Box(
            child = render.Column(
                children = [
                    render.Row(
                        expanded = True,
                        main_align = "center",
                        children = [
                            render.Image(src=GAS_ICON),
                            render.Text(" %d GWEI" % gas_price),
                        ]
                    ),
                    render.Box(
                        height = 4,
                    ),
                    render.Row(
                        expanded = True,
                        main_align = "center",
                        children = [
                            render.Image(src=TIME_ICON),
                            render.Text(" %dh " % hours),
                            render.Text("%dm " % minutes)
                        ]
                    )
                ]
            )
        )
    )