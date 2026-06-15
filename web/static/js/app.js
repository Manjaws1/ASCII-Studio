document.addEventListener('DOMContentLoaded', () => {
    const textInput = document.getElementById('textInput');
    const bannerSelect = document.getElementById('bannerSelect');
    const colorInput = document.getElementById('colorInput');
    const alignSelect = document.getElementById('alignSelect');
    const asciiOutput = document.getElementById('asciiOutput');
    const btnSave = document.getElementById('btnSave');
    const btnNew = document.getElementById('btnNew');
    const btnCopy = document.getElementById('btnCopy');
    const btnDownload = document.getElementById('btnDownload');
    const mobileMenuBtn = document.getElementById('mobileMenuBtn');
    const navLinks = document.getElementById('navLinks');
    const heroSection = document.getElementById('heroSection');
    const btnGetStarted = document.getElementById('btnGetStarted');

    const logo = document.querySelector('.logo');

    // Handle Landing Page logic
    if (localStorage.getItem('ascii_studio_returning') === 'true') {
        if (heroSection) heroSection.classList.add('hidden');
    }

    if (btnGetStarted) {
        btnGetStarted.addEventListener('click', () => {
            localStorage.setItem('ascii_studio_returning', 'true');
            if (heroSection) {
                heroSection.style.opacity = '0';
                setTimeout(() => heroSection.classList.add('hidden'), 500);
            }
            document.getElementById('workspace').scrollIntoView({ behavior: 'smooth' });
            // Optionally focus the text input
            setTimeout(() => textInput.focus(), 600);
        });
    }

    if (logo) {
        logo.addEventListener('click', () => {
            localStorage.removeItem('ascii_studio_returning');
            if (heroSection) {
                heroSection.classList.remove('hidden');
                setTimeout(() => heroSection.style.opacity = '1', 50);
            }
            window.scrollTo({ top: 0, behavior: 'smooth' });
        });
        logo.style.cursor = 'pointer';
    }

    if (mobileMenuBtn && navLinks) {
        mobileMenuBtn.addEventListener('click', () => {
            const isExpanded = navLinks.classList.toggle('show');
            mobileMenuBtn.setAttribute('aria-expanded', isExpanded);
        });
    }

    let debounceTimeout;

    const showToast = (message, type = 'success') => {
        const toastContainer = document.getElementById('toastContainer');
        if (!toastContainer) return;

        const toast = document.createElement('div');
        toast.className = `toast ${type}`;
        toast.innerText = message;
        
        toastContainer.appendChild(toast);
        
        // Trigger reflow to enable animation
        toast.offsetHeight;
        toast.classList.add('show');

        setTimeout(() => {
            toast.classList.remove('show');
            setTimeout(() => {
                toast.remove();
            }, 300); // match transition duration
        }, 3000);
    };

    const generateArt = () => {
        const text = textInput.value;
        const banner = bannerSelect.value;
        const color = colorInput.value;
        const alignment = alignSelect.value;

        if (!text) {
            asciiOutput.innerHTML = 'Type something to generate ASCII art...';
            asciiOutput.style.color = '#fff';
            return;
        }

        fetch('/api/generate', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ text, banner, color, alignment })
        })
        .then(res => res.json())
        .then(data => {
            if (data.error) {
                asciiOutput.innerHTML = `<span style="color: red;">Error: ${data.error}</span>`;
            } else {
                asciiOutput.innerHTML = data.result;
            }
        })
        .catch(err => console.error('Generation error:', err));
    };

    const scheduleGenerate = () => {
        clearTimeout(debounceTimeout);
        debounceTimeout = setTimeout(generateArt, 300);
    };

    textInput.addEventListener('input', scheduleGenerate);
    bannerSelect.addEventListener('change', generateArt);
    colorInput.addEventListener('input', scheduleGenerate);
    alignSelect.addEventListener('change', generateArt);

    btnCopy.addEventListener('click', () => {
        if (!textInput.value) return showToast("Nothing to copy!", "error");
        const textToCopy = asciiOutput.innerText;
        navigator.clipboard.writeText(textToCopy).then(() => {
            showToast('Copied to clipboard!', 'success');
        }).catch(() => {
            showToast('Failed to copy', 'error');
        });
    });

    btnDownload.addEventListener('click', () => {
        const textToDownload = asciiOutput.innerText;
        if (!textInput.value) return showToast("Nothing to download!", "error");
        
        const blob = new Blob([textToDownload], { type: 'text/plain' });
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'ascii_art.txt';
        document.body.appendChild(a);
        a.click();
        window.URL.revokeObjectURL(url);
        document.body.removeChild(a);
    });

    btnNew.addEventListener('click', () => {
        window.history.pushState({}, '', '/');
        textInput.value = '';
        bannerSelect.value = 'standard';
        colorInput.value = '#ffffff';
        alignSelect.value = 'left';
        generateArt();
    });

    btnSave.addEventListener('click', () => {
        const projectData = {
            text: textInput.value,
            banner: bannerSelect.value,
            color: colorInput.value,
            alignment: alignSelect.value
        };

        if (!projectData.text) return showToast("Nothing to save!", "error");

        fetch('/api/save-project', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(projectData)
        })
        .then(res => res.json())
        .then(data => {
            if (data.id) {
                window.history.pushState({}, '', `/?id=${data.id}`);
                showToast('Project saved successfully!', 'success');
            }
        })
        .catch(() => showToast('Failed to save project', 'error'));
    });

    // Check for ID in URL on load
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    if (id) {
        fetch(`/api/load-project/${id}`)
        .then(res => res.json())
        .then(data => {
            if (data.id) {
                textInput.value = data.text;
                bannerSelect.value = data.banner;
                colorInput.value = data.color;
                alignSelect.value = data.alignment;
                generateArt();
            }
        });
    }
});
