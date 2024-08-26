function fb_login() {
  window.location.href = "https://localhost:8080/login";
}

const url = "https://localhost:8080/upload_image";

async function uploadImage(event) {
  event.preventDefault();

  const form = event.currentTarget;
  const formData = new FormData(form);

  console.log(formData.get("fileToUpload"));

  try {
    const response = await fetch(url, {
      method: "POST",
      headers: {},
      body: formData,
    });
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const json = await response.json();
    console.log(json);
  } catch (error) {
    console.error(error.message);
  }
}
