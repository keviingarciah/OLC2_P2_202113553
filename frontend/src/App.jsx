// Vanilla
import './App.css';

// Import React
import React, { useEffect, useState, useRef } from 'react';

// Monaco Editor
import MonacoEditor from 'react-monaco-editor';

// Bootstrap
import 'bootstrap/dist/css/bootstrap.min.css';

// Bootstrap Components
import Table from 'react-bootstrap/Table';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';

// Bootstrap Navbar
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';

// Axios
import axios from 'axios';

//d3 - graphviz 
import * as d3 from 'd3'; // Asegúrate de importar D3
import { graphviz } from 'd3-graphviz';

// CST
import cst from './cst.svg';

// ---------------------------------- Main App ---------------------------------- //
function App() {
  // -----------------Navbar-----------------
  const [selectedFile, setSelectedFile] = useState(null);

  // Abrir archivo
  const handleFileSelect = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = (event) => {
      const newFile = {
        name: file.name,
        content: event.target.result,
      };
      setSelectedFile(newFile);
      setFiles(prevFiles => prevFiles.concat(newFile));
      setCode(newFile.content);
    };
    reader.readAsText(file);
  };

  const handleOpenFileClick = () => {
    document.getElementById('fileInput').click();
  };

  // Cerrar archivo
  const handleCloseFile = () => {
    setSelectedFile(null);
    setCode('');
  };

  // -----------------Archivos-----------------
  const [files, setFiles] = useState([]);

  // Eliminar archivo
  const handleDeleteFile = (index) => {
    const newFiles = [...files];
    const file = files[index];

    // Borrar de la lista
    newFiles.splice(index, 1);
    setFiles(newFiles);

    // Por si estaba abierto
    if (selectedFile != null) {
      if (file.name === selectedFile.name) {
        setSelectedFile(null);
        setCode('');
      }
    }
  };

  // Seleccionar archivo
  const handleSelectFile = (file) => {
    setSelectedFile(file);
    setCode(file.content);
  }

  // Crear archivo
  const handleCreateFile = () => {
    const newFile = {
      name: 'Nuevo Archivo',
      content: '',
    };
    setSelectedFile(newFile);
    setFiles(prevFiles => prevFiles.concat(newFile));
    setCode(newFile.content);

    alert('Archivo creado correctamente.');
  };

  // Guardar archivo
  const handleSaveFile = () => {
    const newFiles = [...files];
    const file = getFileByName();
    file.content = code;
    setFiles(newFiles);
    setSelectedFile(file);

    alert('Archivo guardado correctamente.');
  };

  // Obtener archivo por nombre
  const getFileByName = () => {
    return files.find((file) => file.name === selectedFile.name);
  };

  // -----------------Analizar-----------------
  const [code, setCode] = useState('');
  const [result, setResult] = useState('');

  // Estado para almacenar la tabla de símbolos
  const [symbolTable, setSymbolTable] = useState('');
  // Estado para almacenar la tabla de errores
  const [errorTable, setErrorTable] = useState('');

  // Analizar endpoint
  function analyze() {
    axios
      .post('http://localhost:5000/parser', code, {
        headers: {
          'Content-Type': 'text/plain', // Especifica que el contenido es de tipo texto plano
        },
      })
      .then(function (response) {
        //console.log(response.data.Output);
        //console.log(response.data.Cst);
        //console.log(response.data.SymbolTable);
        //console.log(response.data.ErrorTable);
  
        // Establece setResult con el valor de response.data.output
        setResult(response.data.Output); 
        // Crear tabla de simbolos
        //setSymbolTable(response.data.SymbolTable);
        // Crear tabla de errores
        //setErrorTable(response.data.ErrorTable);
      })
      .catch(function (error) {
        console.log(error);
      });
  }

  // -----------------CST-----------------
  // CST Modal
  const [showCST, setShowCST] = useState(false);

  // Ocultar o Mostrar CST
  const handleCSTClose = () => setShowCST(false);
  const handleCShow = () => setShowCST(true);

    // Analizar endpoint
    function graphCST() {
      handleCShow();  
    }

    const svgContainerRef = useRef(null);
    useEffect(() => {
      if (svgContainerRef.current) {
        const svgContainer = d3.select(svgContainerRef.current);
        const svg = svgContainer.select('svg');
      
        // Configura el comportamiento de zoom y drag
        const zoom = d3.zoom().on('zoom', (event) => {
          svg.attr('transform', event.transform);
        });
      
        svg.call(zoom);
      
        // Centra el contenido SVG en el contenedor
        const width = 1000; // Ancho deseado
        const height = 300; // Alto deseado
        const svgWidth = svg.attr('width');
        const svgHeight = svg.attr('height');
        const scaleX = width / svgWidth;
        const scaleY = height / svgHeight;
        const scale = Math.min(scaleX, scaleY);
      
        svgContainer.call(zoom.transform, d3.zoomIdentity.scale(scale));
      
        // Limpia el zoom cuando se cierra el modal
        return () => {
          svg.on('.zoom', null);
        };
      }
    }, [cst]); // Asegúrate de que este efecto se ejecute cuando cst cambie
    
  // -----------------Tabla de Simbolo-----------------
  const graphRefSTable = useRef(null);

  // TS Modal
  const [showSTable, setShowSTable] = useState(false);

  // Ocultar o Mostrar TS
  const handleSTableClose = () => setShowSTable(false);
  const handleSTableShow = () => setShowSTable(true);

  // Graph SymTable 
  function graphTableSymbol() {    
    handleSTableShow();    
  }

  // -----------------Tabla de Errores-----------------
  const graphRefETable = useRef(null);

  // TE Modal
  const [showETable, setShowETable] = useState(false);

  // Ocultar o Mostrar TE
  const handleETableClose = () => setShowETable(false);
  const handleETableShow = () => setShowETable(true);

  // Graph ErrorTable
  function graphErrorTable() {
    handleETableShow();
  }

  // Render
  return (
    <div className="App">
      <header className="App-header">

        {/* Navbar */}
        <Navbar bg="dark" variant="dark" expand="lg">
          <Container fluid>
            <Navbar.Brand>T-Swift</Navbar.Brand>
            <Navbar.Toggle aria-controls="navbarScroll" />
            <Navbar.Collapse id="navbarScroll">
              <Nav
                className="me-auto my-2 my-lg-0"
                style={{ maxHeight: '100px' }}
                navbarScroll
              >
                <NavDropdown title="Archivo" id="navbarScrollingDropdown">
                  <NavDropdown.Item onClick={handleCreateFile}>
                    Crear Archivo
                  </NavDropdown.Item>
                  <NavDropdown.Divider />
                  <NavDropdown.Item onClick={handleOpenFileClick}>
                    Abrir Archivo
                  </NavDropdown.Item>
                  <NavDropdown.Item onClick={handleSaveFile}>
                    Guardar Archivo
                  </NavDropdown.Item>
                  <NavDropdown.Divider />
                  <NavDropdown.Item onClick={handleCloseFile}>
                    Cerrar
                  </NavDropdown.Item>
                </NavDropdown>

                <NavDropdown title="Reportes" id="navbarScrollingDropdown">
                  <NavDropdown.Item onClick={graphCST}>
                    Árbol CST
                  </NavDropdown.Item>
                  <NavDropdown.Item onClick={graphTableSymbol}>
                    Tabla de Símbolos
                  </NavDropdown.Item>
                  <NavDropdown.Item onClick={graphErrorTable}>
                    Tabla de Errores
                  </NavDropdown.Item>
                </NavDropdown>
              </Nav>

              <Nav className="ml-auto">
                  <Button variant="light" onClick={analyze} className="mr-5">
                    Ejecutar
                  </Button>
              </Nav>    


              <input
                type="file"
                id="fileInput"
                style={{ display: 'none' }}
                onChange={handleFileSelect}
              />
            </Navbar.Collapse>
          </Container>
        </Navbar>

      </header>

      {/* Modal para el CST */}
      <Modal
        show={showCST}
        onHide={handleCSTClose}
        size="lg"
        aria-labelledby="contained-modal-title-vcenter"
        centered
      >
        <Modal.Header closeButton>
        <Modal.Title>CST</Modal.Title>
        </Modal.Header>
        <Modal.Body>
        <div id="svg-container" ref={svgContainerRef}>
          <svg width="1000" height="300">
            <image xlinkHref={cst} width="100%" height="100%" />
          </svg>
        </div>     
        </Modal.Body>
        <Modal.Footer>
          <Button variant="dark" onClick={handleCSTClose}>
            Cerrar
          </Button>
        </Modal.Footer>
      </Modal>

      {/* Modal para la Tabla de Símbolos */}
      <Modal
        show={showSTable}
        onHide={handleSTableClose}
        size="lg"
        aria-labelledby="contained-modal-title-vcenter"
        centered
        onEntered={() => {
          // Llama a la función para renderizar el gráfico aquí
          graphviz(graphRefSTable.current).renderDot(symbolTable);
        }}
      >
        <Modal.Header closeButton>
          <Modal.Title>Tabla de Símbolos</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className="modalSTablegraph">
            <div
              ref={graphRefSTable}
              style={{ width: '700px', display: 'flex', justifyContent: 'center', alignItems: 'center' }}
            ></div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="dark" onClick={handleSTableClose}>
            Cerrar
          </Button>
        </Modal.Footer>
      </Modal>

      {/* Modal para la Tabla de Errores */}
      <Modal
        show={showETable}
        onHide={handleETableClose}
        size="lg"
        aria-labelledby="contained-modal-title-vcenter"
        centered
        onEntered={() => {
          // Llama a la función para renderizar el gráfico aquí
          graphviz(graphRefETable.current).renderDot(errorTable);
        }}
      >
        <Modal.Header closeButton>
          <Modal.Title>Tabla de Errores</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className="modalETablegraph">
            <div
              ref={graphRefETable}
              style={{ width: '700px', display: 'flex', justifyContent: 'center', alignItems: 'center' }}
            ></div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="dark" onClick={handleETableClose}>
            Cerrar
          </Button>
        </Modal.Footer>
      </Modal>

      {/* Tabla y Editores */}
      <Row>
        {/* Tabla de Archivos */}
        <Col md={2} className='files my-4 mr-5'>
          <Table striped bordered hover variant="dark" style={{ width: '105%' }}>
            <thead>
              <tr>
                <th className='text'>Archivos</th>
              </tr>
            </thead>
            <tbody>
              {files.map((file, index) => (
                <tr key={index}>
                  <td>
                    <a
                      href="#"
                      onClick={() => handleSelectFile(file)}
                      className='text-white'
                    >
                      {file.name}
                    </a>
                    <Button
                      variant='link'
                      size='sm'
                      className='text-white text-end'
                      onClick={() => handleDeleteFile(index)}
                      style={{ textAlign: 'right' }}
                    >
                      {"X"}  
                    </Button>
                  </td>
                </tr>
              ))}
            </tbody>
          </Table>
        </Col>

        {/* Editor y Consola */}
        <Col md={10}>
          <div className='editors my-4'>
            <div className='containerE'>
              <div className="editor1 mx-3">
                <MonacoEditor
                  width="900"
                  height="600"
                  language='swift'
                  theme="vs-dark"
                  value={code}
                  options={{ minimap: { enabled: true } }}
                  onChange={setCode}
                />
              </div>
              <div className="editor2">
                <MonacoEditor
                  width="600"
                  height="600"
                  language="swift"
                  theme="vs-dark"
                  value={result}
                  options={{ readOnly: true, minimap: { enabled: false }, lineNumbers: false }}
                />
              </div>
            </div>
          </div>
        </Col>

      </Row>

      {/* Footer */}
      <footer className="blockquote-footer">
        <Card className="bg-dark text-white">
          <Card.Body>
            <Card.Text>
              T-Swift - Organización de Lenguajes y Compiladores 2
            </Card.Text>
          </Card.Body>
        </Card>
      </footer>
    </div>
  );
}

export default App;
