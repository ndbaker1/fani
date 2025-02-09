import { useEffect, useState } from "react"
import { faniClient } from "../lib/fani"
import { FileObject } from "../gen/fani/v1/fani_pb"
import { For } from "@chakra-ui/react"
import { NavLink } from "react-router"
import Image from "../components/Image"

export default function() {
  const [files, setFiles] = useState<FileObject[]>([])

  useEffect(() => {
    faniClient.listStaging({})
      .then(listFilesResponse => setFiles(listFilesResponse.files))
  }, [])

  return (
    <>
      <div style={{ display: "flex", flexWrap: "wrap", gap: 5 }}>
        <For each={files}>
          {(file, i) => (
            <Image
              key={i}
              data={file}
              source={URL.createObjectURL(new Blob([file.thumbnail!]))} />
          )}
        </For>
      </div>
    </>
  )
}
