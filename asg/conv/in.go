package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func InSection(in cobol85.IInSectionContext) *pb.InSection {
	ctx := in.(*cobol85.InSectionContext)
	return &pb.InSection{
		SectionName: SectionName(ctx.SectionName()),
	}
}

func InLibrary(in cobol85.IInLibraryContext) *pb.InLibrary {
	ctx := in.(*cobol85.InLibraryContext)
	return &pb.InLibrary{
		LibraryName: LibraryName(ctx.LibraryName()),
	}
}

func InFile(in cobol85.IInFileContext) *pb.InFile {
	ctx := in.(*cobol85.InFileContext)
	return &pb.InFile{
		FileName: FileName(ctx.FileName()),
	}
}
