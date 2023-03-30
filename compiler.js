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
    } else if (node.type === "function_call") {
        const sourceFunName = node.fun_name.value;
        //const funName = sourceFunName === "if" ? "$if" : sourceFunName;
        const funName = sourceFunName;
        const params = node.parameters.map(generate)
            .join(", ");
        return `${funName}(${params})`;
    } else if (node.type === "function_def") {
        const funName = node.fun_name.value;
        return generateFunction(node.body, node.parameters, funName);
    } else if (node.type === "conditional_ifelse") {
        return generateIfElseConditional(node.comparison, node.if_body, node.else_body)
    } else if (node.type === "conditional_if") {
        return generateIfConditional(node.comparison, node.if_body)
    } else if (node.type === "identifier") {
        return node.value;
    } else if (node.type === "number") {
        return String(node.value);
    } else if (node.type === "string") {
        return node.value;
    } else {
        throw new Error(`Unknown node type: ${node.type}`);
    }
}

function generateFunction(statements, parameters, name = "") {
    const body = statements.map((statement, idx) => {
        const js = generate(statement);
        return js
        // if (idx === statements.length - 1) {
        //     return `return ${js}`;
        // } else {
        //     return js;
        // }
    }).join(";\n") + ";";
    const indentedBody = indent(body);
    const params = parameters.map(generate).join(", ");
    return `function ${name}(${params}) {\n${indentedBody}\n}`;
}

function generateIfElseConditional(comparison, ifBody, elseBody) {
    const ifbody = ifBody.map((statement, idx) => {
        const js = generate(statement);
        return js
    }).join(";\n") + ";";
    const elsebody = elseBody.map((statement, idx) => {
        const js = generate(statement);
        return js
    }).join(";\n") + ";";
    const indentedIfBody = indent(ifbody);
    const indentedElseBody = indent(elsebody);
    return `if (${comparison.expr_one.value} ${comparison.comparison.value} ${comparison.expr_two.value}) {\n${indentedIfBody}\n} else {\n${indentedElseBody}\n}`;
}

function generateIfConditional(comparison, ifBody) {
    const ifbody = ifBody.map((statement, idx) => {
        const js = generate(statement);
        return js
    }).join(";\n") + ";";
    const indentedIfBody = indent(ifbody);
    return `if (${comparison.expr_one.value} ${comparison.comparison.value} ${comparison.expr_two.value}) {\n${indentedIfBody}\n}`;
}

function indent(string) {
    return string.split("\n").map(line => "\t" + line).join("\n");
}


main().catch(err => console.log(err.stack));