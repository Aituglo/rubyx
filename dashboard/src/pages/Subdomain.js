import React, { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import {
  createSubdomain,
  deleteSubdomain,
  getSubdomains,
  updateSubdomain,
} from "../actions/subdomain";
import PageTitle from "../components/Typography/PageTitle";
import { TrashIcon, EditIcon } from "../icons";
import {
  Table,
  TableHeader,
  TableCell,
  TableBody,
  TableRow,
  TableFooter,
  TableContainer,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Button,
  Pagination,
  Label,
  Select,
  Input,
} from "@windmill/react-ui";
import { getPrograms } from "../actions/program";

function Subdomain() {
  const dispatch = useDispatch();
  const subdomainState = useSelector((state) => state.subdomain);
  const programState = useSelector((state) => state.program);
  const [url, setUrl] = useState("");
  const [program, setProgram] = useState();
  const [editId, setEditId] = useState(0);
  const [editMode, setEditMode] = useState(false);

  const [isModalOpen, setIsModalOpen] = useState(false);

  function openModal() {
    setEditMode(false);
    setIsModalOpen(true);
  }

  function closeModal() {
    setIsModalOpen(false);
  }

  useEffect(() => {
    dispatch(getPrograms());
    dispatch(getSubdomains());
  }, []);

  const handleCreateSubdomain = () => {
    setIsModalOpen(false);

    dispatch(
      createSubdomain(parseInt(program), url)
    );

    setUrl("");
    setProgram(0);
  };

  const handleUpdateSubdomain = () => {
    setIsModalOpen(false);
    setEditMode(false);

    dispatch(
      updateSubdomain(
        parseInt(editId),
        parseInt(program),
        url
      )
    );

    setUrl("");
    setProgram(0);
    setEditId(0);
  };

  const handleDeleteSubdomain = (id) => {
    dispatch(deleteSubdomain(id));
  };

  const handleEditSubdomain = (subdomain) => {
    setUrl(subdomain.url);
    setProgram(subdomain.program_id);
    setEditId(subdomain.id);
    setEditMode(true);
    setIsModalOpen(true);
  };

  const [pageTable, setPageTable] = useState(1);

  const [dataTable, setDataTable] = useState([]);

  // pagination setup
  const resultsPerPage = 10;
  const totalResults = subdomainState.subdomains ? subdomainState.subdomains.length : 0;

  function onPageChangeTable(p) {
    setPageTable(p);
  }

  const getProgramName = (id) => {
    if (programState && programState.programs) {
      var potential = programState.programs.find(
        (item) => item.id == parseInt(id)
      );
      if (potential) {
        return potential.name;
      } else {
        return "";
      }
    }
  };

  useEffect(() => {
    setDataTable(
      subdomainState.subdomains &&
        subdomainState.subdomains.slice(
          (pageTable - 1) * resultsPerPage,
          pageTable * resultsPerPage
        )
    );
  }, [pageTable, subdomainState]);

  return (
    <>
      <PageTitle>Subdomains</PageTitle>

      <div className="px-4 py-3 mb-8 bg-white rounded-lg shadow-md dark:bg-gray-800">
        <div className="py-3">
          <Button onClick={openModal}>Add a Subdomain</Button>
        </div>

        {totalResults > 0 && (
          <TableContainer className="mb-8">
            <Table>
              <TableHeader>
                <tr>
                  <TableCell>Subdomain</TableCell>
                  <TableCell>Program</TableCell>
                  <TableCell>Title</TableCell>
                  <TableCell>Status Code</TableCell>
                  <TableCell>Content Length</TableCell>
                  <TableCell>Actions</TableCell>
                </tr>
              </TableHeader>
              <TableBody>
                {dataTable &&
                  dataTable.map((key, i) => (
                    <TableRow key={i}>
                      <TableCell>
                        <span className="text-sm">{key.url}</span>
                      </TableCell>
                      <TableCell>
                        <span className="text-sm">
                          {getProgramName(key.program_id)}
                        </span>
                      </TableCell>
                      <TableCell>
                        <span className="text-sm">{key.title}</span>
                      </TableCell>
                      <TableCell>
                        <span className="text-sm">{key.status_code}</span>
                      </TableCell>
                      <TableCell>
                        <span className="text-sm">{key.content_length}</span>
                      </TableCell>
                      <TableCell>
                        <div className="flex items-center space-x-4">
                          <Button layout="link" size="icon" aria-label="Delete">
                            <TrashIcon
                              onClick={() => handleDeleteSubdomain(key.id)}
                              className="w-5 h-5"
                              aria-hidden="true"
                            />
                          </Button>
                          <Button layout="link" size="icon" aria-label="Delete">
                            <EditIcon
                              onClick={() => handleEditSubdomain(key)}
                              className="w-5 h-5"
                              aria-hidden="true"
                            />
                          </Button>
                        </div>
                      </TableCell>
                    </TableRow>
                  ))}
              </TableBody>
            </Table>
            <TableFooter>
              <Pagination
                totalResults={totalResults}
                resultsPerPage={resultsPerPage}
                onChange={onPageChangeTable}
                label="Navigation"
              />
            </TableFooter>
          </TableContainer>
        )}

        <Modal isOpen={isModalOpen} onClose={closeModal}>
          <ModalHeader>Add a Subdomain</ModalHeader>
          <ModalBody>
            {programState && programState.programs && (
              <Label className="pt-5">
                <span>Program</span>
                <Select
                  value={program}
                  onChange={(e) => setProgram(e.target.value)}
                  className="mt-1"
                >
                  <option value="">Select a Program</option>
                  {programState.programs.map((item) => (
                    <option value={item.id}>{item.name}</option>
                  ))}
                </Select>
              </Label>
            )}
            <Label className="pt-5">
              <span>Subdomain</span>
              <Input
                className="mt-1"
                placeholder=""
                value={url}
                onChange={(e) => setUrl(e.target.value)}
              />
            </Label>
          </ModalBody>
          <ModalFooter>
            <div className="hidden sm:block">
              <Button layout="outline" onClick={closeModal}>
                Cancel
              </Button>
            </div>
            <div className="hidden sm:block">
              {editMode ? (
                <Button onClick={() => handleUpdateSubdomain()}>Update</Button>
              ) : (
                <Button onClick={() => handleCreateSubdomain()}>Add</Button>
              )}
            </div>
            <div className="block w-full sm:hidden">
              <Button block size="large" layout="outline" onClick={closeModal}>
                Cancel
              </Button>
            </div>
            <div className="block w-full sm:hidden">
              {editMode ? (
                <Button
                  block
                  size="large"
                  onClick={() => handleUpdateSubdomain()}
                >
                  Update
                </Button>
              ) : (
                <Button
                  block
                  size="large"
                  onClick={() => handleCreateSubdomain()}
                >
                  Add
                </Button>
              )}
            </div>
          </ModalFooter>
        </Modal>
      </div>
    </>
  );
}

export default Subdomain;
