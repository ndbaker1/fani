import { useEffect, useState } from "react"
import { faniClient } from "../lib/fani"
import { FileObject } from "../gen/fani/v1/fani_pb"

export default function() {
  const [files, setFiles] = useState<FileObject[]>([])

  useEffect(() => {
    faniClient.listStaging({})
      .then(listFilesResponse => setFiles(listFilesResponse.files))
  }, [])

  return (
    <>
      {files.map(file => (
        <div>
          {JSON.stringify(file.attributes)}
        </div>
      ))}
    </>
  )
}
