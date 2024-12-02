document.addEventListener('DOMContentLoaded', function() {
    const copyButtons = document.querySelectorAll('.copy-btn');
  
    copyButtons.forEach(button => {
      button.addEventListener('click', function() {
        const codeBlock = this.previousElementSibling;
        const text = codeBlock.innerText;
  
        navigator.clipboard.writeText(text).then(() => {
          alert('Code copied to clipboard!');
        }).catch(err => {
          console.error('Failed to copy text: ', err);
        });
      });
    });
  });