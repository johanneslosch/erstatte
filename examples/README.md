# Examples

This directory contains example files to demonstrate how to use `erstatte`.

## Files

- **`example-erstatte.json`** - Example configuration file showing how to define replacement operations
- **`testfile.ts`** - Example TypeScript file that can be modified using the settings

## Usage

From the project root directory:

```bash
# Run with the example configuration
./erstatte --settings examples/example-erstatte.json

# Or from within the examples directory
cd examples
../erstatte --settings example-erstatte.json
```

## Configuration Format

The settings file supports the following structure:

```json
{
  "boolean": [
    {
      "filePath": "./path/to/file.ts",
      "field": "fieldName",
      "value": "true|false"
    }
  ]
}
```

### Fields

- **`filePath`**: Path to the file to modify (supports .json, .ts, .properties)
- **`field`**: Name of the field/property to replace
- **`value`**: New value to set (strings, booleans, numbers)

## Supported File Types

- **JSON** (`.json`) - Replaces values in JSON objects
- **TypeScript** (`.ts`) - Replaces values in TypeScript configuration objects
- **Properties** (`.properties`) - Replaces key-value pairs in properties files
