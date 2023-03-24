const fs = require("fs").promises;
const { readFileSync } = require("fs");
const path = require("path");
const RUNTIME = readFileSync(path.join(__dirname, "runtime.js")).toString();

async function main() {
    const filename = process.argv[2];
    if (!filename) {
        console.log("Please provide a file name.");
        return;
    }
    const astCode = (await fs.readFile(filename)).toString();
    const ast = JSON.parse(astCode);
    const jsCode = generate(ast);
    const baseDir = path.dirname(filename);
    const baseName = path.basename(filename, ".x.ast");
    const jsFilename = path.join(baseDir, `${baseName}.js`);
    await fs.writeFile(jsFilename, jsCode);
    console.log(`Wrote ${jsFilename}.`);
}

const declaredVars = new Map();

function generate(node) {
    if (node.type === "program") {
        return node.body.map(generate).join(";\n") + ";\n" + RUNTIME;
    } else if (node.type === "assignment") {
        const varName = node.var_name.value;
        const value = generate(node.value);
        declaredVars.set(varName, true)
        return `var ${varName} = ${value}`;
    } else if (node.type === "reassignment") { 
        const varName = node.var_name.value;
        const value = generate(node.value);
        if (declaredVars.get(varName)) {
            return `${varName} = ${value}`;
        } else {
            throw new Error('this variable is missing initial declaration. please use var keyword to declare this variable.')
        }
    }else if (node.type === "identifier") {
        return node.value;
    } else if (node.type === "number") {
        return String(node.value);
    } else if (node.type === "string") {
        return node.value;
    } else {
        throw new Error(`Unknown node type: ${node.type}`);
    }
}


main().catch(err => console.log(err.stack));