// Tab switching
document.querySelectorAll('.tab-btn').forEach(btn => {
    btn.addEventListener('click', () => {
        const tab = btn.dataset.tab;
        
        // Update buttons
        document.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));
        btn.classList.add('active');
        
        // Update content
        document.querySelectorAll('.tab-content').forEach(c => c.classList.remove('active'));
        document.querySelector(`.tab-content[data-tab="${tab}"]`).classList.add('active');
    });
});

// Simple SakeraLang Interpreter (Client-side simulator)
class SakeraInterpreter {
    constructor() {
        this.variables = {};
        this.output = [];
    }
    
    run(code) {
        this.output = [];
        this.variables = {};
        
        try {
            const lines = code.split('\n');
            
            for (let line of lines) {
                line = line.trim();
                
                // Skip empty lines and comments
                if (!line || line.startsWith('//')) continue;
                
                this.executeLine(line);
            }
            
            return this.output.join('\n');
        } catch (error) {
            return `Error: ${error.message}`;
        }
    }
    
    executeLine(line) {
        // Remove trailing comment
        line = line.split('//')[0].trim();
        
        // Variable assignment: sango x = value
        if (line.startsWith('sango ')) {
            const match = line.match(/sango\s+(\w+)\s*=\s*(.+)/);
            if (match) {
                const varName = match[1];
                const value = this.evaluate(match[2]);
                this.variables[varName] = value;
            }
        }
        // Print: toles value
        else if (line.startsWith('toles ')) {
            const expr = line.substring(6).trim();
            const value = this.evaluate(expr);
            this.output.push(String(value));
        }
    }
    
    evaluate(expr) {
        expr = expr.trim();
        
        // String literal
        if (expr.startsWith('"') && expr.endsWith('"')) {
            return expr.slice(1, -1);
        }
        
        // Number
        if (!isNaN(expr)) {
            return parseFloat(expr);
        }
        
        // Boolean
        if (expr === 'bender') return true;
        if (expr === 'sala') return false;
        
        // Variable
        if (this.variables.hasOwnProperty(expr)) {
            return this.variables[expr];
        }
        
        // Simple arithmetic
        if (expr.includes('+')) {
            const parts = expr.split('+').map(p => this.evaluate(p.trim()));
            return parts.reduce((a, b) => {
                if (typeof a === 'string' || typeof b === 'string') {
                    return String(a) + String(b);
                }
                return a + b;
            });
        }
        
        if (expr.includes('-')) {
            const parts = expr.split('-').map(p => this.evaluate(p.trim()));
            return parts.reduce((a, b) => a - b);
        }
        
        if (expr.includes('*')) {
            const parts = expr.split('*').map(p => this.evaluate(p.trim()));
            return parts.reduce((a, b) => a * b);
        }
        
        if (expr.includes('/')) {
            const parts = expr.split('/').map(p => this.evaluate(p.trim()));
            return parts.reduce((a, b) => a / b);
        }
        
        return expr;
    }
}

// Run button handler
const interpreter = new SakeraInterpreter();

document.getElementById('run-btn').addEventListener('click', () => {
    const code = document.getElementById('code-editor').value;
    const output = interpreter.run(code);
    document.getElementById('output').textContent = output || '(no output)';
});

// Smooth scrolling
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            });
        }
    });
});

// Run default code on page load
window.addEventListener('load', () => {
    document.getElementById('run-btn').click();
});
