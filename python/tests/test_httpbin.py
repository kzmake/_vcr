import pytest
import requests

@pytest.mark.default_cassette("httpbin_status_200.yaml")
@pytest.mark.vcr
def test_httpbin_get_status_200():
    assert requests.get("https://httpbin.org/status/200").status_code == 200

@pytest.mark.default_cassette("httpbin_fakestatus_200.yaml")
@pytest.mark.vcr
def test_httpbin_get_status_200_fake():
    assert requests.get("https://httpbin.org/fakestatus/200").status_code == 200
